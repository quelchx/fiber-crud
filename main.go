package main

import (
	"os"

	"github.com/gofiber/fiber"
	"github.com/quelchx/fiber-crud/initializers"
	"github.com/quelchx/fiber-crud/services"
)

func init() {
	initializers.LoadEnv()
	initializers.DatabaseClient()
	initializers.SyncDatabase()
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	services.UserRoutes(app)
	services.PostRoutes(app)

	port := os.Getenv("PORT")
	app.Listen(port)
}
