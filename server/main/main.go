package main

import (
	"github.com/LukeTarr/Tckt/api"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	// set production mode to false until we collect env vars
	prod := false
	port := os.Getenv("PORT")
	// if we have $PORT set, we can assume we're in production
	if port != "" {
		prod = true
	}

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
