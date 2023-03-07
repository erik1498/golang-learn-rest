package book

import (
	"learn-rest/config"
	bookHandler "learn-rest/internal/handlers/book"

	jwtware "github.com/gofiber/jwt/v3"

	"github.com/gofiber/fiber/v2"
)

func SetupBookRoutes(router fiber.Router) {
	book := router.Group("/book")

	book.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Config("JWT_SECRET")),
	}))

	book.Get("/", bookHandler.GetAllBook)
	book.Get("/:bookId", bookHandler.GetBook)
	book.Post("/create", bookHandler.CreateBook)
	book.Put("/:bookId", bookHandler.UpdateBook)
	book.Delete("/:bookId", bookHandler.DeleteBook)
}
