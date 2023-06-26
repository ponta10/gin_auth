// ./app/models/user.go
package models

import (
	"gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "time"
)

var JwtKey = []byte("your_secret_key")

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

type User struct {
    gorm.Model
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
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

func (user *User) GenerateToken() (string, error) {
    expirationTime := time.Now().Add(5 * time.Minute)

    claims := &Claims{
        Username: user.Username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(JwtKey)
}
