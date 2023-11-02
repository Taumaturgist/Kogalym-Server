package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"kogalym-backend/helpers"
	"kogalym-backend/models"
	"net/http"
)

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

type SignedResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

type Client struct {
	Login          string
	HashedPassword string
}

type LoginData struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var loginParams LoginData
	if err := c.ShouldBindJSON(&loginParams); err != nil {
		helpers.ValidateError(c, err)
		return
	}

	client, err := models.GetClientByLogin(loginParams.Login)

	helpers.CheckErr(err)
	if client.Login == "" || !CheckPasswordHash(loginParams.Password, client.HashedPassword) {
		c.JSON(http.StatusBadRequest, UnsignedResponse{
			Message: "Bad credentials",
		})
		return
	}

	token, err := generateToken(client.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SignedResponse{
		Token:   token,
		Message: "logged in",
	})

	return
}

func JwtTokenCheck(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := parseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "bad jwt token",
		})
		return
	}

	_, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
			Message: "unable to parse claims",
		})
		return
	}
	c.Next()
}

func PrivateACLCheck(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := parseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "bad jwt token",
		})
		return
	}

	claims, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
			Message: "unable to parse claims",
		})
		return
	}

	claimedUID, OK := claims["user"].(string)
	if !OK {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "no user property in claims",
		})
		return
	}

	uid := c.Param("uid")
	if claimedUID != uid {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "token uid does not match resource uid",
		})
		return
	}

	c.Next()
}
