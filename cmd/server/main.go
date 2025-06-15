package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Vikram-D16/MyPath-be/db"
	"github.com/Vikram-D16/MyPath-be/handlers"
	"github.com/Vikram-D16/MyPath-be/middlewares"
	"github.com/Vikram-D16/MyPath-be/models"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	dbInstance := db.DB

	dbInstance.AutoMigrate(&models.User{})

	router := gin.Default()

	// Authenticate request
	router.Use(func(c *gin.Context) {
		if c.FullPath() != "/register" && c.FullPath() != "/login" {
			middlewares.AuthMiddleware()(c)
			if c.IsAborted() {
				return
			}
		}
		c.Next()
	})

	// Define routes
	router.POST("/register", handlers.RegisterHandler)
	router.POST("/login", handlers.LoginHandler)
	router.DELETE("/account", handlers.DeleteUserHandler)

	router.GET("/home", handlers.HomeHandler)
	router.GET("/users", handlers.GetAllUsersHandler)

	router.POST("/organizations", handlers.CreateOrganizationHandler)
	router.GET("/organizations", handlers.GetAllOrganizationsHandler)

	router.POST("/groups", handlers.CreateGroupHandler)
	router.GET("/groups", handlers.GetAllGroupsHandler)
	router.DELETE("/groups/:id", handlers.DeleteGroupHandler)

	router.POST("/group-member", handlers.CreateGroupMemberHandler)

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
