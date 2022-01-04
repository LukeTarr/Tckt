package api

import (
	"github.com/gofiber/fiber/v2"
)

func HandleRegister(c *fiber.Ctx) error {
	// Create a Register form instance to hold the body into
	r := new(Register)

	// Populate the instance with data from the body
	err := c.BodyParser(r)
	if err != nil {
		return err
	}

	// If a field is left blank, return an error
	if r.Email == "" || r.Password == "" || r.PasswordConfirm == "" {
		return c.JSON(fiber.Map{"error": "Fill out all inputs"})
	}

	// If the passwords do not match, return an error
	if r.Password != r.PasswordConfirm {
		return c.JSON(fiber.Map{"error": "Passwords do not match"})
	}

	// TODO: check against database, make sure no user with same email exists, if not push this user in

	// return a success message
	return c.JSON("success")
}

func HandleLogin(c *fiber.Ctx) error {
	l := new(Login)

	err := c.BodyParser(l)
	if err != nil {
		return err
	}

	if l.Email == "" || l.Password == "" {
		return c.JSON(fiber.Map{"error": "Fill out all inputs"})
	}

	// TODO: grab user from database, check their passwords

	return c.JSON("success")
}
