package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	app := fiber.New() // create a new Fiber instance

	// Create a new endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!") // send text
	})

	// Look for $PORT (env var for Heroku) to listen on or use 4000 for local dev
	port := os.Getenv("PORT")
	if port == "" {
		err := app.Listen(":4000")
		if err != nil {
			return
		}
	} else {
		err := app.Listen(":" + port)
		if err != nil {
			return
		}
	}
}
