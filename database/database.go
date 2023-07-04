package database

import (
	"fmt"
	"learn-rest/config"
	"learn-rest/internal/book/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitialDB() {
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("DB_USER"),
		config.Config("DB_PASS"),
		config.Config("DB_HOST"),
		config.Config("DB_PORT"),
		config.Config("DB_NAME"))
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to Initial DB")
	}

	log.Println("Initial DB Success")
	DB.AutoMigrate(&models.Book{})
	log.Print("Migrate DB Success")

}
