package auth

import (
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"kogalym-backend/helpers"
	"net/http"
	"os"
	"strings"
)

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
		"title": "Main website",
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
		setErrorJson(c, "Parameters can't be empty")
		return
	}

	// Check for username and password match, usually from a database
	if login != os.Getenv("ADMIN_LOGIN") || !CheckPasswordHash(password, os.Getenv("ADMIN_PASSWORD")) {
		setErrorJson(c, "Неправильный логин или пароль")
		return
	}

	// Save the username in the session
	SetSessionValue(c, sessionUserKey, login)

	c.JSON(http.StatusOK, gin.H{"error": "Failed to save session"})
}

// todo
//func WebLogout(c *gin.Context) {
//	session := sessions.Default(c)
//	user := session.Get(sessionUserKey)
//
//	if user == nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
//		return
//	}
//
//	session.Delete(sessionUserKey)
//	if err := session.Save(); err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
//}
