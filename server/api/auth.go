package api

import (
	"github.com/LukeTarr/Tckt/database"
	"github.com/LukeTarr/Tckt/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HandleRegister(c *fiber.Ctx) error {
	// Create a Register form instance to hold the body into
	r := Register{}

	// Populate the instance with data from the body
	err := c.BodyParser(&r)
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

	// Check against database, make sure no user with same email exists, if not push this user in
	var u models.User
	res := database.DB.Where("email = ?", r.Email).First(&u)
	if res.RowsAffected != 0 {
		return c.JSON(fiber.Map{"error": "User with that email already exists"})
	}

	// Grab the encryption key from env var and hash the password with it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}

	// make a user with data from the post body
	u = models.User{
		Role:     "User",
		Email:    r.Email,
		Password: string(hashedPassword),
	}

	//save the user to the database
	database.DB.Save(&u)

	// return a success message
	return c.JSON("success")
}

func HandleLogin(c *fiber.Ctx) error {
	l := Login{}

	// Parse the input into a Login instance
	err := c.BodyParser(&l)
	if err != nil {
		return err
	}

	// Return an error if not all fields are filled out
	if l.Email == "" || l.Password == "" {
		return c.JSON(fiber.Map{"error": "Fill out all inputs"})
	}

	// Create a user and store the database call into the user
	var u models.User
	res := database.DB.Where("email = ?", l.Email).First(&u)

	// If the database doesn't find a user matching that email, return an error
	if res.RowsAffected == 0 {
		return c.JSON(fiber.Map{"error": "No user with email: " + l.Email})
	}

	// Compare the returned user's hashed password to the password in the input
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(l.Password))
	if err != nil {
		return c.JSON(fiber.Map{"error": "Incorrect password"})
	}

	// TODO: generate a JWT to send to the user

	// return a success message
	return c.JSON("success")
}
