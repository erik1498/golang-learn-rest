package routers

import (
	"learn-rest/internal/middleware/jwt/routers"
	book "learn-rest/internal/routers/book"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	book.SetupBookRoutes(api)
	routers.SetupRoutesJwt(api)
}
