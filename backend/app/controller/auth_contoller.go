package controller

import (
	"Kogalym/backend/app/service"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	LoginPage(c *gin.Context)
	WebLogin(c *gin.Context)
	WebLogout(c *gin.Context)
}

type AuthControllerImpl struct {
	svc service.AuthService
}

func (a AuthControllerImpl) LoginPage(c *gin.Context) {
	a.svc.LoginPage(c)
}

func (a AuthControllerImpl) WebLogin(c *gin.Context) {
	a.svc.WebLogin(c)
}

func (a AuthControllerImpl) WebLogout(c *gin.Context) {
	a.svc.WebLogout(c)
}

func AuthControllerInit(authService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		svc: authService,
	}
}
