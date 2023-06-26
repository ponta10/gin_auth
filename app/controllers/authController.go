// ./app/controllers/userController.go
package controllers

import (
	"fmt"
	"gin_auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterEndpoint(c *gin.Context, db *gorm.DB) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := user.HashPassword(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
        return
    }

    if err := models.CreateUser(db, &user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "User created successfully"})
}

func LoginEndpoint(c *gin.Context, db *gorm.DB) {
    var user, foundUser models.User
    fmt.Println(user, foundUser)
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // user の値を foundUser にコピー
    foundUser = user

    fmt.Println(user, foundUser)

    if err := models.FindUserByUsername(db, &foundUser); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding user"})
        return
    }

    if err := foundUser.CheckPassword(user.Password); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
        return
    }

    token, err := foundUser.GenerateToken()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

