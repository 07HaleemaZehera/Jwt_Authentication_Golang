package controllers

import (
	"JWT_AUTHENTICATION/database"
	"JWT_AUTHENTICATION/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("hello world")
}


func Register(c *fiber.Ctx) error{
	fmt.Println("Recieved a registration request")

	// map[KeyType]ValueType
	 var data map[string]string
	 if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	 }
	//  fiber.Map is a shorthand for map[string]interface{}

// }

 // Check if the email already exists
 var existingUser models.User
 if err := database.DB.Where("email = ?", data["email"]).First(&existingUser).Error; err == nil {
	 return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		 "error": "Email already exists",
	 })
 }
 hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: hashedPassword,
	}
	// Insert user into database
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
	})
}
