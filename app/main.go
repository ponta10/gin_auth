// ./main.go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "gin_auth/app/controllers"
    "gin_auth/app/models"
    "gin_auth/app/middlewares"
)

func main() {
    db, err := gorm.Open("postgres", "host=localhost user=gorm dbname=gorm password=gorm sslmode=disable")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    db.AutoMigrate(&models.User{}, &models.Post{})

    r := gin.Default()

    r.POST("/register", func(c *gin.Context) { controllers.RegisterEndpoint(c, db) })
    r.POST("/login", func(c *gin.Context) { controllers.LoginEndpoint(c, db) })

    r.POST("/posts", middlewares.AuthMiddleware(), func(c *gin.Context) { controllers.CreatePostEndpoint(c, db) })
    
    r.Run(":8080")
}
