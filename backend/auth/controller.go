package auth

import (
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
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
	login := c.PostForm("login")
	password := c.PostForm("password")

	// Validate form input
	if strings.Trim(login, " ") == "" || strings.Trim(password, " ") == "" {
		setError(c, "Parameters can't be empty", "/login")
		return
	}

	// Check for username and password match, usually from a database
	if login != os.Getenv("ADMIN_LOGIN") || !CheckPasswordHash(password, os.Getenv("ADMIN_PASSWORD")) {
		setError(c, "Authentication failed", "/login")
		return
	}

	// Save the username in the session
	SetSessionValue(c, sessionUserKey, login)

	c.Redirect(301, "/")
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
