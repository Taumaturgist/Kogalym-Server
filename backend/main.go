package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"kogalym-backend/auth"
	"kogalym-backend/models"
	"net/http"
	"os"
)

//go:embed templates/*
var f embed.FS

func main() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	logFile, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	router := gin.Default()

	templates := template.Must(template.New("").ParseFS(f, "templates/*.html"))
	router.SetHTMLTemplate(templates)

	models.ConnectDatabase()

	//log.Fatal(auth.HashPassword("dron"))

	// Auth
	authRouter := router.Group("")
	{
		authRouter.POST("/login", auth.Login)
	}

	// API
	v1 := router.Group("/api")
	v1.Use(auth.JwtTokenCheck)
	{
		v1.GET("/", getPersons)
		v1.Use(auth.PrivateACLCheck).GET("/:uid", getPersons)
	}

	// WEB
	web := router.Group("")
	{
		web.GET("", home)
		// Статичные файлы фронта
		web.StaticFS("/public", http.Dir("templates/public"))
	}

	router.Run(":8080")
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}

func getPersons(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "Something"})
}
