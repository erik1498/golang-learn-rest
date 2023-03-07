package handlers

import (
	"learn-rest/config"
	"learn-rest/internal/helper"
	"learn-rest/internal/middleware/jwt/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {
	var login models.Login

	err := c.BodyParser(&login)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": http.StatusBadRequest, "message": "Review your input", "data": err})
	}

	errValidation := helper.ValidateStruct(&login)
	if errValidation != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http.StatusBadRequest, "message": "Error Validation", "validation": errValidation})
	}

	if login.Username != "John" || login.Password != "Doe" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"status":  http.StatusUnauthorized,
			"message": "Unauthorized",
		})
	}

	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Minute * 10).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": http.StatusInternalServerError, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"token": t})
}

// func restricted(c *fiber.Ctx) error {
// 	user := c.Locals("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	name := claims["name"].(string)
// 	return c.SendString("Welcome " + name)
// }
