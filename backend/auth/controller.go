package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"kogalym-backend/helpers"
	"net/http"
	"os"
	"strings"
)

type LoginData struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func WebAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := GetSessionUser(c)

		if user == nil {
			setError(c, "Требуется вход в систему", "/login")
			return
		}

		c.Next()
	}
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"csrf":  csrf.GetToken(c),
		"error": getError(c),
	})
}

func WebLogin(c *gin.Context) {
	var loginParams LoginData
	if err := c.ShouldBindJSON(&loginParams); err != nil {
		helpers.ValidateError(c, err)
		return
	}
	login := loginParams.Login
	password := loginParams.Password

	if strings.Trim(login, " ") == "" || strings.Trim(password, " ") == "" {
		setErrorJson(c, "Заполните логин и пароль")
		return
	}

	// Check for username and password match, usually from a database
	if login != os.Getenv("ADMIN_LOGIN") || password != os.Getenv("ADMIN_PASSWORD") {
		setErrorJson(c, "Неправильный логин или пароль")
		return
	}

	// Save the username in the session
	SetSessionValue(c, sessionUserKey, login)
}

func WebLogout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(sessionUserKey)

	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}

	session.Delete(sessionUserKey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
