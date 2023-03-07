package bookHandler

import (
	"learn-rest/database"
	request "learn-rest/internal/dto/request"
	"learn-rest/internal/dto/response"
	"learn-rest/internal/helper"
	"learn-rest/internal/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllBook(c *fiber.Ctx) error {
	db := database.DB
	var book []response.BookResponse

	db.Raw("SELECT * FROM `books`").Scan(&book)
	// db.Find(&book)

	if len(book) == 0 {
		return c.Status(http.StatusNotFound).JSON(
			fiber.Map{
				"status":  http.StatusNotFound,
				"message": "Data Is Empty",
				"data":    nil,
			},
		)
	}

	return c.Status(http.StatusOK).JSON(
		fiber.Map{
			"status":  http.StatusOK,
			"message": "Success",
			"data":    book,
		},
	)
}

func CreateBook(c *fiber.Ctx) error {
	db := database.DB
	book := new(models.Book)

	err := c.BodyParser(book)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": http.StatusBadRequest, "message": "Review your input", "data": err})
	}

	book.ID = uuid.New()

	err = db.Create(&book).Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http.StatusBadRequest, "message": "Error create book", "data": nil})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"status": http.StatusCreated, "message": "Book created success", "data": book})
}

func GetBook(c *fiber.Ctx) error {
	db := database.DB
	var book models.Book

	id := c.Params("bookId")
	db.Find(&book, "id = ?", id)

	if book.ID == uuid.Nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": "Data Not Found",
			"data":    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": http.StatusOK, "message": "Success", "data": book})
}

func UpdateBook(c *fiber.Ctx) error {
	db := database.DB
	var book models.Book

	id := c.Params("bookId")
	db.Find(&book, "id = ?", id)

	if book.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": http.StatusNotFound, "message": "Data not found", "data": nil})
	}
	var updateBook request.UpdateBook
	err := c.BodyParser(&updateBook)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": http.StatusInternalServerError, "message": "Review your input", "data": nil})
	}

	errValidation := helper.ValidateStruct(&updateBook)
	if errValidation != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http.StatusBadRequest, "message": "Error Validation", "validation": errValidation})
	}

	book.Title = updateBook.Title
	db.Save(&book)

	return c.Status(http.StatusNoContent).JSON(fiber.Map{"status": http.StatusNoContent, "message": "Success", "data": nil})
}

func DeleteBook(c *fiber.Ctx) error {
	db := database.DB
	var book models.Book

	id := c.Params("bookId")
	db.Find(&book, "id = ?", id)

	if book.ID == uuid.Nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"status": http.StatusNotFound, "message": "Data not found", "data": nil})
	}

	err := db.Delete(&book, "id = ?", id).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http.StatusBadRequest, "message": "Failed to delete book", "data": nil})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"status": http.StatusOK, "message": "Success", "data": nil})
}
