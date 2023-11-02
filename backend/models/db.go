package models

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"kogalym-backend/helpers"
	"log"
	"os"
)

var DB *sql.DB

func ConnectDatabase() error {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbPath := os.Getenv("DB_PATH")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		helpers.CheckErr(err)
		return err
	}

	DB = db
	return nil
}
