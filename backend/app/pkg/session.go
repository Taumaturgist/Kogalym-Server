package pkg

import (
	"Kogalym/backend/app/constant"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const sessionUserKey = "user"

func SetSession(c *gin.Context, value any) {
	session := sessions.Default(c)
	session.Set(sessionUserKey, value)

	if err := session.Save(); err != nil {
		log.Error("Ошибка в установки сесии")
		PanicException(constant.UnknownError)
		return
	}
}

func DeleteSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(sessionUserKey)
}

func GetSessionUser(c *gin.Context) any {
	session := sessions.Default(c)
	return session.Get(sessionUserKey)
}
