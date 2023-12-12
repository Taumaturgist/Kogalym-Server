package middleware

import (
	"Kogalym/backend/app/constant"
	"Kogalym/backend/app/pkg"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func WebAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := pkg.GetSessionUser(c)

		if user == nil {
			log.Error("Требуется вход в систему")
			pkg.PanicException(constant.InvalidRequest)
			c.Redirect(http.StatusSeeOther, "/login")
			return
		}

		c.Next()
	}
}
