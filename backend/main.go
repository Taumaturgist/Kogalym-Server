package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"kogalym-backend/auth"
	"kogalym-backend/business"
	"kogalym-backend/graphParser"
	"kogalym-backend/models"
	"net/http"
	"os"
)

//go:embed templates/*
var f embed.FS

func main() {
	//fmt.Println(auth.HashPassword(os.Getenv("ADMIN_PASSWORD")))

	setupLogging()

	router := gin.Default()

	createRoutesForStaticFiles(router)

	models.ConnectDatabase()

	graphParser.ParseData()

	apiRoutes(router)

	webRoutes(router)

	router.Run(":80")
}

func apiRoutes(router *gin.Engine) {
	// API
	v1 := router.Group("/api")

	// Auth
	v1.POST("/login", auth.Login)

	v1.Use(auth.JwtTokenCheck)
	{
		// todo getSettings
		v1.GET("", getPersons)
		v1.Use(auth.PrivateACLCheck).GET("/:uid", getPersons)
	}
}

func webRoutes(router *gin.Engine) {
	web := router.Group("")

	auth.SetupSessionStore(web)
	auth.SetupCsrfTokens(web)

	// Generate CSRF token
	web.Use(auth.CsrfCheckMiddleware())

	// Auth
	{
		web.GET("/login", auth.LoginPage)
		web.POST("/login", auth.WebLogin)
	}

	authenticatedWeb := web.Group("")
	authenticatedWeb.Use(auth.WebAuthMiddleware())
	{
		authenticatedWeb.GET("", business.Home)
		// todo получение настроек

		// todo требует доработки, перенаправления на главную страницу
		//authenticatedWeb.POST("/logout", auth.WebLogout)
		// todo изменение настроек
	}
}

func setupLogging() {
	logFile, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
}

func createRoutesForStaticFiles(router *gin.Engine) {
	templates := template.Must(template.New("").ParseFS(f, "templates/*.html"))
	router.SetHTMLTemplate(templates)
	router.StaticFS("/public", http.Dir("templates/public"))
}

func getPersons(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "Something"})
}
