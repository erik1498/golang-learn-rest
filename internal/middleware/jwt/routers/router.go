package routers

import (
	"learn-rest/internal/middleware/jwt/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutesJwt(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/login", handlers.Login)
}
