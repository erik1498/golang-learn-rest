package book

import (
	"learn-rest/config"
	"learn-rest/internal/book/handlers"

	jwtware "github.com/gofiber/jwt/v3"

	"github.com/gofiber/fiber/v2"
)

func SetupBookRoutes(router fiber.Router) {
	book := router.Group("/book")

	book.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Config("JWT_SECRET")),
	}))

	book.Get("/", handlers.GetAllBook)
	book.Get("/:bookId", handlers.GetBook)
	book.Post("/create", handlers.CreateBook)
	book.Put("/:bookId", handlers.UpdateBook)
	book.Delete("/:bookId", handlers.DeleteBook)
}
