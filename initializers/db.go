package initializers

import (
	"log"
	"os"

	"github.com/TwiN/go-color"
	model "github.com/quelchx/fiber-crud/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// gorm database client for the application
// var GormClient *gorm.DB

// connects to the database
func DatabaseClient() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	model.GormClient, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(color.Ize(color.Red, "Error connecting to database"))
	}

	log.Println(color.Ize(color.Cyan, "Connected to database"))

}
