package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates/*
var f embed.FS

func main() {
	router := gin.Default()

	templates := template.Must(template.New("").ParseFS(f, "templates/*.html"))
	router.SetHTMLTemplate(templates)

	// API
	v1 := router.Group("/api")
	{
		v1.GET("/", getPersons)
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

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
