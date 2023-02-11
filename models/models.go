package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// replica of the gorm.Model struct but lowercase fields
type Base struct {
	// ID        uint           `gorm:"primarykey" json:"id"`
	// unique id
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Post struct {
	Base
	Title   string `json:"title"`
	Content string `json:"content"`
}

type User struct {
	Base
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}
