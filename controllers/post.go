package controllers

import (
	"github.com/google/uuid"
	model "github.com/quelchx/fiber-crud/models"
)

func CreatePost() (model.Post, error) {
	var post model.Post
	post.UUID = uuid.New()

	tx := model.GormClient.Create(&post)

	if tx.Error != nil {
		return post, tx.Error
	}

	return post, nil
}

func GetAllPosts() ([]model.Post, error) {
	var posts []model.Post

	tx := model.GormClient.Find(&posts)

	if tx.Error != nil {
		return posts, tx.Error
	}

	return posts, nil
}

func GetPostById(id string) (model.Post, error) {
	var post model.Post

	tx := model.GormClient.Find(&post, id)

	if tx.Error != nil {
		return post, tx.Error
	}

	return post, nil
}

func DeletePostById(id string) error {
	var post model.Post

	tx := model.GormClient.Unscoped().Delete(&post, id)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func UpdatePostById(id string, post model.Post) (model.Post, error) {
	tx := model.GormClient.Model(&post).Where("id = ?", id).Updates(post)

	if tx.Error != nil {
		return post, tx.Error
	}

	return post, nil
}
