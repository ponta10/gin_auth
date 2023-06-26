// ./app/controllers/userController.go
package controllers

import (
	"gin_auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterEndpoint(c *gin.Context, db *gorm.DB) {
    var user models.User
    // c.ShouldBindJSON(&user)を用いてクライアントから送られてきたJSONデータをuserオブジェクトにバインド（解析）します
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 入力されたパスワードのハッシュ化
    if err := user.HashPassword(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
        return
    }

    // ユーザーの登録
    if err := models.CreateUser(db, &user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
        return
    }

    // 200番を返す
    c.JSON(http.StatusOK, gin.H{"status": "User created successfully"})
}

func LoginEndpoint(c *gin.Context, db *gorm.DB) {
    var user, foundUser models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // user の値を foundUser にコピー
    foundUser = user

    // 入力された名前と一致するuserの検索
    if err := models.FindUserByUsername(db, &foundUser); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding user"})
        return
    }

    // パスワードのチェック
    if err := foundUser.CheckPassword(user.Password); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
        return
    }

    // tokenの発行
    token, err := foundUser.GenerateToken()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

