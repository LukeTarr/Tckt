package database

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB gorm.DB

func init() {
	// check to see if on dev machine, load .env if so
	if os.Getenv("DATABASE_URL") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			return
		}
	}

	dsn := os.Getenv("DATABASE_URL")

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB!")
	}
	DB = *conn
}
