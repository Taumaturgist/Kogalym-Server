package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//go:embed templates/*
var f embed.FS

func main() {
	router := gin.Default()

	templ := template.Must(template.New("").ParseFS(f, "templates/*.html"))
	router.SetHTMLTemplate(templ)

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.StaticFS("/public", http.Dir("templates/public"))

	router.Run(":8080")
}
