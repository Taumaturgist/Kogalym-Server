package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	csrf "github.com/utrack/gin-csrf"
	"net/http"
	"os"
)

func CsrfCheckMiddleware() gin.HandlerFunc {
	return csrf.Middleware(csrf.Options{
		Secret: os.Getenv("CSRF_SECRET"),
		ErrorFunc: func(c *gin.Context) {
			log.Error("CSRF token mismatch")
			c.JSON(http.StatusBadRequest, gin.H{"error": [1]string{"CSRF token mismatch"}})
			//pkg.PanicException(constant.InvalidRequest)
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
		},
	})
}
