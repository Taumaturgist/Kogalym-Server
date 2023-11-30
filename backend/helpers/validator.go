package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ApiError struct {
	Field string
	Msg   string
}

func ValidateError(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	out := make(map[string]string)

	if errors.As(err, &ve) {
		for _, fe := range ve {
			out[fe.Field()] = fe.Tag()
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": out})
	}
}

func WebValidateError(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	var out []string

	if errors.As(err, &ve) {
		for _, fe := range ve {
			out = append(out, fe.Field()+` `+fe.Tag())
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": out})
	}
}
