package main

import (
	"gin-demo/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := os.Getenv("DATABASE_URL")
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to auto migrate database:", err)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {
		models.GetUsers(c, db)
	})
	router.POST("/users", func(c *gin.Context) {
		models.CreateUser(c, db)
	})
	router.GET("/users/:id", func(c *gin.Context) {
		models.GetUserByID(c, db)
	})
	router.PUT("/users/:id", func(c *gin.Context) {
		models.UpdateUser(c, db)
	})
	router.DELETE("/users/:id", func(c *gin.Context) {
		models.DeleteUser(c, db)
	})

	router.GET("/auth/google/login", models.HandleGoogleLogin)
	router.GET("/auth/google/callback", func(c *gin.Context) {
		models.HandleGoogleCallback(c)
	})

	return router
}

func main() {
	initDB()
	router := setupRouter()

	log.Println("Server is starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
