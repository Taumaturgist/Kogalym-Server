package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"os"
)

func SetupCsrfTokens(group *gin.RouterGroup) {
	store := cookie.NewStore([]byte("store"))
	group.Use(sessions.Sessions("session", store))
}

func CsrfCheckMiddleware() gin.HandlerFunc {
	return csrf.Middleware(csrf.Options{
		Secret: os.Getenv("CSRF_SECRET"),
		ErrorFunc: func(c *gin.Context) {
			setError(c, "CSRF token mismatch", "/login")
		},
	})
}
