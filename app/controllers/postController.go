// ./app/controllers/postController.go
package controllers

import (
	"gin_auth/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreatePostEndpoint(c *gin.Context, db *gorm.DB) {
    // c.Get("userId")を用いてリクエストのcontextからuserIdを取得します。このuserIdは、前段階のミドルウェアであるAuthMiddleware()によって設定されたもので、現在認証されているユーザーのIDを表します。
    userId, exists := c.Get("userId")
    if !exists {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
        return
    }
    var post models.Post

    // c.ShouldBindJSON(&post)を用いてリクエストボディのJSONをPost型のオブジェクトpostにバインドします。
    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // postのUserIDフィールドに、contextから取得したuserIdをセットします。
    post.UserID = userId.(uint)

    // models.CreatePost(db, &post)を用いてpostをデータベースに保存します。
    if err := models.CreatePost(db, &post); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating post"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "Post created successfully"})
}

func GetUserPosts(c *gin.Context, db *gorm.DB) {
    // c.Get("userId")を用いてリクエストのcontextからuserIdを取得します。
    userId, exists := c.Get("userId")
    if !exists {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
        return
    }

    // 証されたユーザー（userId）の投稿を全て取得します。
    user, err := models.GetUserWithPosts(db, userId.(uint))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user posts"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"posts": user.Posts})
}