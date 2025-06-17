package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Vikram-D16/MyPath-be/db"
	"github.com/Vikram-D16/MyPath-be/handlers"
	"github.com/Vikram-D16/MyPath-be/middlewares"
	"github.com/Vikram-D16/MyPath-be/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	dbInstance := db.DB
	dbInstance.AutoMigrate(&models.User{})

	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")

	// Auth middleware for protected routes
	api.Use(func(c *gin.Context) {
		if c.FullPath() != "/api/register" && c.FullPath() != "/api/login" {
			middlewares.AuthMiddleware()(c)
			if c.IsAborted() {
				return
			}
		}
		c.Next()
	})

	api.POST("/register", handlers.RegisterHandler)
	api.POST("/login", handlers.LoginHandler)
	api.DELETE("/account", handlers.DeleteUserHandler)

	api.GET("/home", handlers.HomeHandler)
	api.GET("/users", handlers.GetAllUsersHandler)

	api.POST("/organizations", handlers.CreateOrganizationHandler)
	api.GET("/organizations", handlers.GetAllOrganizationsHandler)

	api.POST("/groups", handlers.CreateGroupHandler)
	api.GET("/groups", handlers.GetAllGroupsHandler)
	api.DELETE("/groups/:id", handlers.DeleteGroupHandler)

	api.POST("/group-member", handlers.CreateGroupMemberHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	fmt.Printf("Server running on http://localhost:%s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
