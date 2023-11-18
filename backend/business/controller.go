package business

import (
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"net/http"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"csrf":  csrf.GetToken(c),
		"title": "Добро пожаловать на главную страницу",
	})
}
