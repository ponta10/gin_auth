// ./app/models/post.go
package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
    UserID uint `json:"user_id"`
}

func CreatePost(db *gorm.DB, post *Post) error {
	return db.Create(post).Error
}
