package repository

import (
	"learn-rest/database"
	"learn-rest/internal/book/models"
)

func GetAllBook() []models.BookResponse {
	db := database.DB
	var book []models.BookResponse
	db.Raw("SELECT * FROM `books`").Scan(&book)
	return book
}

func CreateBook(book models.Book) error {
	db := database.DB
	return db.Create(&book).Error
}
