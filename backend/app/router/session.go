package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func SetupSession() gin.HandlerFunc {
	var sessionStore = cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))

	sessionLifetime, _ := strconv.Atoi(os.Getenv("SESSION_LIFETIME"))
	sessionStore.Options(sessions.Options{MaxAge: sessionLifetime})

	return sessions.Sessions("session", sessionStore)
}
