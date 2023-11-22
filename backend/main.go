package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"html/template"
	"io"
	"kogalym-backend/auth"
	"kogalym-backend/business"
	"log"
	"net/http"
	"os"
)

//go:embed templates/*
var f embed.FS

func main() {
	loadEnv()
	setupLogging()

	router := gin.Default()

	createRoutesForStaticFiles(router)

	webRoutes(router)

	fmt.Println(os.Getenv("PORT"))

	router.Run(":" + os.Getenv("PORT"))
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
		authenticatedWeb.POST("/logout", auth.WebLogout)
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

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
