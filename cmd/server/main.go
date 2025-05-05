package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Vikram-D16/MyPath-be/db"
	"github.com/Vikram-D16/MyPath-be/handlers"
	"github.com/Vikram-D16/MyPath-be/models"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	dbInstance := db.DB

	dbInstance.AutoMigrate(&models.User{})

	router := gin.Default()

	// Define routes
	router.GET("/home", handlers.HomeHandler)
	router.POST("/register", handlers.RegisterHandler)
	router.GET("/users", handlers.GetAllUsersHandler)
	router.POST("/login", handlers.LoginHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// Start the HTTP server
	fmt.Printf("Server running on http://localhost:%s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
