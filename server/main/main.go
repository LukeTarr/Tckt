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

	port := os.Getenv("PORT")
	if port == "" {
		app.Listen(":4000")
	} else {
		app.Listen(":" + port)
	}
}
