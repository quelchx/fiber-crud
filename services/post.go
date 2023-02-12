package services

import (
	"github.com/gofiber/fiber"

	"github.com/quelchx/fiber-crud/controllers"
	model "github.com/quelchx/fiber-crud/models"
)

// create a post
func CreatePost(c *fiber.Ctx) {
	post, err := controllers.CreatePost()

	if err != nil {
		c.Status(503).JSON(fiber.Map{
			"message": "Error creating post",
		})
		return
	}

	c.JSON(post)
}

func GetPosts(c *fiber.Ctx) {
	posts, err := controllers.GetAllPosts()

	if err != nil {
		c.Status(503).JSON(fiber.Map{
			"message": "Error getting posts",
		})
		return
	}

	c.JSON(posts)
}

// get post by id
func GetPostById(c *fiber.Ctx) {
	id := c.Params("id")

	post, err := controllers.GetPostById(id)

	if err != nil {
		c.Status(503).JSON(fiber.Map{
			"message": "Error getting post",
		})
		return
	}

	c.JSON(post)
}

// delete a post by id
func DeletePost(c *fiber.Ctx) {
	id := c.Params("id")

	err := controllers.DeletePostById(id)

	if err != nil {
		c.Status(503).JSON(fiber.Map{
			"message": "Error deleting post",
		})
		return
	}

	c.JSON(fiber.Map{
		"message": "Post deleted",
	})
}

// update a post by id
func UpdatePost(c *fiber.Ctx) {
	id := c.Params("id")
	var post model.Post

	if err := c.BodyParser(&post); err != nil {
		c.Status(503).JSON(fiber.Map{
			"message": "Error parsing post",
		})
		return
	}

	post, err := controllers.UpdatePostById(id, post)

	if err != nil {
		c.Status(503).JSON(fiber.Map{
			"message": "Error updating post",
		})
		return
	}

	c.JSON(post)

}

// post routes
func PostRouter(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/posts", GetPosts)
	api.Get("/post/:id", GetPostById)
	api.Post("/post", CreatePost)
	api.Delete("/post/:id", DeletePost)
	api.Patch("/post/:id", UpdatePost)
}
