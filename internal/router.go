package internal

import (
	"learn-rest/internal/book"
	"learn-rest/internal/middleware/jwt/routers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	book.SetupBookRoutes(api)
	routers.SetupRoutesJwt(api)
}
