// ./models/user.go
package models

import (
    "github.com/jinzhu/gorm"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    gorm.Model
    Username string `json:"username"`
    Password string `json:"password"`
}

func (user *User) HashPassword() error {
    hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user.Password = string(hash)
    return nil
}

func (user *User) CheckPassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func CreateUser(db *gorm.DB, user *User) error {
    return db.Create(user).Error
}

func FindUserByUsername(db *gorm.DB, user *User) error {
    return db.Where("username = ?", user.Username).First(user).Error
}
