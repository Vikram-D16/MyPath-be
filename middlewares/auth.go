package middlewares

import (
	"net/http"
	"strings"

	"github.com/Vikram-D16/MyPath-be/db"
	"github.com/Vikram-D16/MyPath-be/models"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		var userToken models.UserToken
		err := db.DB.
			Where("token = ? AND is_active = ?", token, true).
			First(&userToken).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Attach userId to the context so it’s available in handlers
		c.Set("userId", userToken.UserID)

		c.Next()
	}
}
