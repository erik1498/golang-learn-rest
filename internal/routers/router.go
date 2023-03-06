package routers

import (
	bookRoutes "learn-rest/internal/routers/book"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	bookRoutes.SetupBookRoutes(api)
}
