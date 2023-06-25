package controllers

import (
    "github.com/gin-gonic/gin"
    "gin_auth/app/models"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetPostsEndpoint(c *gin.Context, db *gorm.DB) {
    var posts []models.Post
    if err := models.GetAllPosts(db, &posts); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving posts"})
        return
    }
    c.JSON(http.StatusOK, posts)
}

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

    c.JSON(http.StatusOK, post)
}