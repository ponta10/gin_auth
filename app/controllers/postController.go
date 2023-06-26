// ./app/controllers/postController.go
package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gin_auth/app/models"
	"net/http"
)

func CreatePostEndpoint(c *gin.Context, db *gorm.DB) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreatePost(db, &post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Post created successfully"})
}
