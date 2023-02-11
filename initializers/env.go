package initializers

import (
	"log"

	"github.com/TwiN/go-color"

	"github.com/joho/godotenv"
)

// loads environment variables from .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(color.Ize(color.RedBackground, "Error loading .env file"))
	}

	log.Println(color.Ize(color.Purple, "Loaded environment variables"))
}
