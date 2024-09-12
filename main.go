package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"models"
)

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable"
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

	router.GET("/users", models.GetUsers)
	router.POST("/users", models.CreateUser)
	router.GET("/users/:id", models.GetUserById)
	router.PUT("/users/:id", models.UpdateUser)
	router.DELETE("/users/:id", models.DeleteUser)

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