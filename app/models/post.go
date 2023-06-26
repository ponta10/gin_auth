// ./app/models/post.go
package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}

func CreatePost(db *gorm.DB, post *Post) error {
	return db.Create(post).Error
}
