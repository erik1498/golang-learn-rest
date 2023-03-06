package bookRoutes

import (
	bookHandler "learn-rest/internal/handlers/book"

	"github.com/gofiber/fiber/v2"
)

func SetupBookRoutes(router fiber.Router) {
	book := router.Group("/book")

	book.Get("/", bookHandler.GetAllBook)
	book.Get("/:bookId", bookHandler.GetBook)
	book.Post("/create", bookHandler.CreateBook)
	book.Put("/:bookId", bookHandler.UpdateBook)
	book.Delete("/:bookId", bookHandler.DeleteBook)
}
