// ./main.go
package main

import (
	"fmt"
	"gin_auth/controllers"
	"gin_auth/middlewares"
	"gin_auth/models"

	"github.com/gin-gonic/gin"
    gormigrate "github.com/go-gormigrate/gormigrate/v2"
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


	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20230601",
			Migrate: func(tx *gorm.DB) error {
				err := tx.Migrator().DropTable("users", "posts")
				if err != nil {
					return err
				}
				err = tx.AutoMigrate(&models.User{}, &models.Post{})
				if err != nil {
					return err
				}
				// Seed data here if necessary
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users", "posts")
			},
		},
	})

	if err := m.Migrate(); err != nil {
		fmt.Printf("Could not migrate: %v", err)
	}

	r := gin.Default()

	r.POST("/register", func(c *gin.Context) { controllers.RegisterEndpoint(c, db) })
	r.POST("/login", func(c *gin.Context) { controllers.LoginEndpoint(c, db) })

	r.POST("/posts", middlewares.AuthMiddleware(), func(c *gin.Context) { controllers.CreatePostEndpoint(c, db) })
    r.GET("/myposts", middlewares.AuthMiddleware(), func(c *gin.Context) { controllers.GetUserPosts(c, db) })
    
	r.Run(":8080")
}
