package middleware

import (
	"backend-api/config"
	"backend-api/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		var idToken models.PersonalAccessToken
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak ditemukan!"})
			c.Abort()
			return
		}
		if err := config.DB.Where("token = ?", strings.TrimPrefix(token, "Bearer ")).First(&idToken).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid!"})
			c.Abort()
			return

		}
		c.Next()
	}
}
