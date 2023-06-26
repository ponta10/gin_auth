// ./app/models/post.go
package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}

func CreatePost(db *gorm.DB, post *Post) error {
	return db.Create(post).Error
}
