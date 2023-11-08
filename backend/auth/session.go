package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

const sessionUserKey = "user"

var sessionStore = cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
var sessionName = "session"

func SetSessionValue(c *gin.Context, key any, value any) {
	session := sessions.Default(c)
	session.Set(key, value)

	if err := session.Save(); err != nil {
		setError(c, "Failed to save session", "/")
		return
	}
}

func GetSessionValueAndDelete(c *gin.Context, key any) string {
	session := sessions.Default(c)
	value := session.Get(key)
	session.Delete(key)

	if err := session.Save(); err != nil {
		setError(c, "Failed to save session", "/")
	}

	if value == nil {
		return ""
	}

	return value.(string)
}

func GetSessionUser(c *gin.Context) any {
	session := sessions.Default(c)
	return session.Get(sessionUserKey)
}

func SetupSessionStore(router *gin.RouterGroup) {
	sessionLifetime, _ := strconv.Atoi(os.Getenv("SESSION_LIFETIME"))

	sessionStore.Options(sessions.Options{MaxAge: sessionLifetime})
	router.Use(sessions.Sessions(sessionName, sessionStore))
}
