// ./main.go
package main

import (
	"fmt"
	"gin_auth/controllers"
	"gin_auth/middlewares"
	"gin_auth/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "user=gorm password=gorm dbname=gorm host=db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	db.AutoMigrate(&models.User{}, &models.Post{})

	r := gin.Default()

	r.POST("/register", func(c *gin.Context) { controllers.RegisterEndpoint(c, db) })
	r.POST("/login", func(c *gin.Context) { controllers.LoginEndpoint(c, db) })

	r.POST("/posts", middlewares.AuthMiddleware(), func(c *gin.Context) { controllers.CreatePostEndpoint(c, db) })
    r.GET("/myposts", middlewares.AuthMiddleware(), func(c *gin.Context) { controllers.GetUserPosts(c, db) })
    
	r.Run(":8080")
}
