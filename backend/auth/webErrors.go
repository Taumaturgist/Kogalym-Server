package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setErrorJson(c *gin.Context, errorMessage string) {
	errorArray := []string{errorMessage}
	c.JSON(http.StatusBadRequest, gin.H{"error": errorArray})
}

func setError(c *gin.Context, errorMessage string, route string) {
	SetSessionValue(c, "error", errorMessage)

	if "" != route {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	c.Redirect(http.StatusSeeOther, route)
}

func getError(c *gin.Context) string {
	return GetSessionValueAndDelete(c, "error")
}
