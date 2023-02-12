package services

import (
	"os"
	"time"

	"github.com/gofiber/fiber"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	model "github.com/quelchx/fiber-crud/models"
	"github.com/quelchx/fiber-crud/validator"
	"golang.org/x/crypto/bcrypt"
)

// checks if the user's email already exists

// sign up a new user (hashes password)
func SignUp(c *fiber.Ctx) {
	var user model.User
	user.UUID = uuid.New()

	if err := c.BodyParser(&user); err != nil {
		c.Status(503).JSON(fiber.Map{
			"message": "Error parsing user",
		})
		return
	}

	// check if the user already exists
	existingUser := validator.IsValidUser(user.Email)

	if existingUser.Email != "" {
		c.Status(409).JSON(fiber.Map{
			"message": "User already has an account",
		})
		return
	}

	// hash the password
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.Status(503).JSON(fiber.Map{
			"message": "Error hashing password",
		})
		return
	}

	user.Password = string(password)
	model.GormClient.Create(&user)

	c.JSON(user)
}

func Login(c *fiber.Ctx) {
	// get email and password from the body
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		c.Status(503).JSON(fiber.Map{
			"message": "Error parsing user",
		})
		return
	}

	existingUser := validator.IsValidUser(user.Email)

	if existingUser.Email == "" {
		c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))

	if err != nil {
		c.Status(401).JSON(fiber.Map{
			"message": "Incorrect password",
		})
		return
	}

	signature := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": existingUser.ID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	token, err := signature.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.Status(503).JSON(fiber.Map{
			"message": "Error signing token",
		})
		return
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 1),
		HTTPOnly: true,
		SameSite: "lax",
	})

	c.JSON(fiber.Map{
		"token": token,
	})
}

func Logout(c *fiber.Ctx) {
	if c.Cookies("token") == "" {
		c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
		return
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		SameSite: "lax",
	})

	c.JSON(fiber.Map{
		"message": "Logged out",
	})
}

// get the current user
func GetCurrentUser(c *fiber.Ctx) {
	cookie := c.Cookies("token")

	token, err := jwt.Parse(
		cookie,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, nil
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

	if err != nil {
		c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
		return
	}

	var user model.User
	model.GormClient.Where("id = ?", claims["sub"]).First(&user)

	c.JSON(user)
}

func UserRouter(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Post("/signup", SignUp)
	api.Post("/login", Login)
	api.Post("/logout", Logout)
	api.Get("/user", GetCurrentUser)
}
