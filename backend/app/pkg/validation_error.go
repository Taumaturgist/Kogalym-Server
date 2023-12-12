package pkg

import (
	"Kogalym/backend/app/constant"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func ValidationError(c *gin.Context, data []string) {
	log.Error("Validation exception")
	c.JSON(http.StatusBadRequest, BuildResponse_(
		constant.InvalidRequest.GetResponseStatus(),
		"Validation exception",
		data,
	))
	c.Abort()
}
