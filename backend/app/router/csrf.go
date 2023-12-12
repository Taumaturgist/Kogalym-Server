package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupCsrfTokens() gin.HandlerFunc {
	store := cookie.NewStore([]byte("store"))
	return sessions.Sessions("session", store)
}
