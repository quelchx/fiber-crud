package services

import (
	"github.com/gofiber/fiber"
	"github.com/google/uuid"

	"github.com/quelchx/fiber-crud/initializers"
	model "github.com/quelchx/fiber-crud/models"
)

// create a post
func CreatePost(c *fiber.Ctx) {
	var post model.Post
	post.ID = uuid.New()

	if err := c.BodyParser(&post); err != nil {
		c.Status(503).JSON(fiber.Map{
			"message": "Error parsing post",
		})
		return
	}

	initializers.GormClient.Create(&post)

	c.JSON(post)
}

// get all the posts
func GetPosts(c *fiber.Ctx) {
	var posts []model.Post

	if err := initializers.GormClient.Find(&posts).Error; err != nil {
		c.Status(404).JSON(fiber.Map{
			"message": "Posts not found",
		})
		return
	}

	c.JSON(posts)
}

// get post by id
func GetPostById(c *fiber.Ctx) {
	var post model.Post

	id := c.Params("id")

	// find the post by id and check for errors
	initializers.GormClient.Find(&post, id)

	// if there is no post return
	if post.Title == "" {
		c.Status(404).JSON(fiber.Map{
			"message": "Post not found",
		})
		return
	}

	c.JSON(post)
}

// delete a post by id
func DeletePost(c *fiber.Ctx) {
	var post model.Post

	id := c.Params("id")

	// find the post by id and check for errors
	initializers.GormClient.Find(&post, id)

	// if there is no post return
	if post.Title == "" {
		c.Status(404).JSON(fiber.Map{
			"message": "Post not found",
		})
		return
	}

	initializers.GormClient.Delete(&post)

	c.JSON(fiber.Map{
		"message": "Post deleted",
		"post":    post,
	})
}

// update a post by id
func UpdatePost(c *fiber.Ctx) {
	var post model.Post

	id := c.Params("id")

	// check to see if the post exists
	initializers.GormClient.Find(&post, id)

	// if there is no post return
	if post.Title == "" {
		c.Status(404).JSON(fiber.Map{
			"message": "Post not found",
		})
		return
	}

	// parse the body of the request
	if err := c.BodyParser(&post); err != nil {
		c.Status(503).JSON(fiber.Map{
			"message": "Error parsing post",
		})
		return
	}

	// update the post
	initializers.GormClient.Save(&post)

	c.JSON(post)

}

// post routes
func PostRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/posts", GetPosts)
	api.Get("/post/:id", GetPostById)
	api.Post("/post", CreatePost)
	api.Delete("/post/:id", DeletePost)
	api.Patch("/post/:id", UpdatePost)
}
