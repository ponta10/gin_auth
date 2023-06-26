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

// db.Create(post) の部分で新しい投稿のレコードをデータベースに保存しています
func CreatePost(db *gorm.DB, post *Post) error {
	return db.Create(post).Error
}

func GetUserWithPosts(db *gorm.DB, userID uint) (*User, error) {
    var user User
    // reloadメソッドはGORMのEager Loadingを実現するためのメソッドで、この場合"Posts"という引数はUserモデルにおける関連付けられたPostモデルをEager Loadingすることを指示しています。
    // つまり、ユーザー情報を取得する際に同時にそのユーザーの投稿も取得します。
    if err := db.Preload("Posts").First(&user, userID).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
