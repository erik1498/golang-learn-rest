package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	error := godotenv.Load(".env")

	if error != nil {
		log.Println("Error Load", key, "From Environtment")
	}

	return os.Getenv(key)
}
