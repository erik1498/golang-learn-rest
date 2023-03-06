package main

import (
	"learn-rest/database"
	"learn-rest/internal/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.InitialDB()
	routers.SetupRoutes(app)
	app.Listen(":8080")
}
