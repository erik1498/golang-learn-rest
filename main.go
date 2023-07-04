package main

import (
	"learn-rest/database"
	"learn-rest/internal"

	// jwtware "github.com/gofiber/jwt/v3"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	database.InitialDB()
	internal.SetupRoutes(app)

	app.Use(cors.New(cors.ConfigDefault))
	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: []byte(config.Config("JWT_SECRET")),
	// }))

	app.Listen(":8080")
}
