// ./app/controllers/postController.go
package controllers

import (
	"gin_auth/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreatePostEndpoint(c *gin.Context, db *gorm.DB) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.UserID = userId.(uint)

	if err := models.CreatePost(db, &post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Post created successfully"})
}
