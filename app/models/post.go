package models

import (
    "github.com/jinzhu/gorm"
)

type Post struct {
    gorm.Model
    Title   string `json:"title"`
    Content string `json:"content"`
    UserID  uint   `json:"user_id"`
}

func GetAllPosts(db *gorm.DB, posts *[]Post) error {
    return db.Find(posts).Error
}

func CreatePost(db *gorm.DB, post *Post) error {
    return db.Create(post).Error
}

