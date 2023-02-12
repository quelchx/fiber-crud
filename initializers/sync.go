package initializers

import (
	"log"

	"github.com/TwiN/go-color"
	model "github.com/quelchx/fiber-crud/models"
)

// will run any migrations on the database
func SyncDatabase() {
	migrate := model.GormClient.AutoMigrate(&model.Post{}, &model.User{})

	if migrate != nil {
		log.Fatal(color.Ize(color.Red, "Error migrating database"))
	}

	log.Println(color.Ize(color.Green, "Database migrated successfully"))
}
