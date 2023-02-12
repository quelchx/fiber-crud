package validator

import (
	model "github.com/quelchx/fiber-crud/models"
)

func IsValidUser(email string) model.User {
	var user model.User
	model.GormClient.Where("email = ?", email).First(&user)
	return user
}
