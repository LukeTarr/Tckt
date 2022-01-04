package main

import (
	"fmt"
	"github.com/LukeTarr/Tckt/api"
	"github.com/gofiber/fiber/v2"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// set production mode to false until we collect env vars
	prod := false
	port := os.Getenv("PORT")
	// if we have $PORT set, we can assume we're in production
	if port != "" {
		prod = true
	}

	// if dev, load our env from .env
	if !prod {
		err := godotenv.Load(".env")
		if err != nil {
			panic("Missing .env file!")
		}
	}

	// Grab the Database credentials from env var
	dsn := os.Getenv("DATABASE_URL")

	// Connect to the database with our credentials
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Couldn't connect to DB!")
	}

	fmt.Println(db)

	// create a new Fiber instance
	app := fiber.New()

	// attach  handlers from api package
	auth := app.Group("auth")

	auth.Post("/register", api.HandleRegister)
	auth.Post("/login", api.HandleLogin)

	// Listen on $PORT for prod or :4000 for dev
	if prod {
		app.Listen(":" + port)
	} else {
		app.Listen(":4000")
	}

}
