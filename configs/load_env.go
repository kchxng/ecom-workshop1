package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	log.Println("Loading .env file successfully...")
}

func GetENV(k string) string {
	return os.Getenv(k)
}
