package main

import (
	"Kogalym/backend/app/router"
	"Kogalym/backend/config"
	"embed"
	"github.com/joho/godotenv"
	"os"
)

//go:embed templates/*
var embeddedFiles embed.FS

func init() {
	godotenv.Load()
	config.InitLog()
	config.InitLog()
}

func main() {
	port := os.Getenv("PORT")

	init := config.Init()
	app := router.Init(init, embeddedFiles)

	app.Run(":" + port)
}
