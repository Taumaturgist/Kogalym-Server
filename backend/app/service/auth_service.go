package service

import (
	"Kogalym/backend/app/constant"
	"Kogalym/backend/app/pkg"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	csrf "github.com/utrack/gin-csrf"
	"net/http"
	"os"
	"strings"
)

type AuthService interface {
	LoginPage(c *gin.Context)
	WebLogin(c *gin.Context)
	WebLogout(c *gin.Context)
}

type AuthServiceImpl struct {
}

type LoginData struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a AuthServiceImpl) LoginPage(c *gin.Context) {
	defer pkg.PanicHandler(c)

	c.HTML(http.StatusOK, "login.html", gin.H{
		"csrf": csrf.GetToken(c),
	})
}

func (a AuthServiceImpl) WebLogin(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var loginParams LoginData
	if err := c.ShouldBindJSON(&loginParams); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}
	login := strings.Trim(loginParams.Login, " ")
	password := strings.Trim(loginParams.Password, " ")

	// Check for username and password match, usually from a database
	if login != os.Getenv("ADMIN_LOGIN") || password != os.Getenv("ADMIN_PASSWORD") {
		pkg.ValidationError(c, []string{"Неправильный логин или пароль"})
	}

	// Save the username in the session
	pkg.SetSession(c, login)
}

func (a AuthServiceImpl) WebLogout(c *gin.Context) {
	defer pkg.PanicHandler(c)

	user := pkg.GetSessionUser(c)

	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}

	pkg.DeleteSession(c)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func AuthServiceInit() *AuthServiceImpl {
	return &AuthServiceImpl{}
}
