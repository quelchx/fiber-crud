package main

import (
	"log"

	"github.com/quelchx/fiber-crud/initializers"
	model "github.com/quelchx/fiber-crud/models"
)

func init() {
	initializers.LoadEnv()
	initializers.DatabaseClient()
}

func main() {
	migrate := initializers.GormClient.AutoMigrate(&model.Post{}, &model.User{})

	if migrate != nil {
		log.Fatal("Error migrating database")
	}

	log.Println("Migration successful")
}
