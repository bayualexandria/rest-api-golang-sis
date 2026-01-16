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
		var user models.User
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak ditemukan!","status":401})
			c.Abort()
			return
		}
		if err := config.DB.Where("token = ?", strings.TrimPrefix(token, "Bearer ")).First(&idToken).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak valid!","status":401})
			c.Abort()
			return

		}
		if err := config.DB.Where("status_id = ?", "4").First(&user).Error; err != nil {
			c.JSON(403, gin.H{"message": "User ini tidak memiliki akses!","status":403})
			c.Abort()
			return

		}
		c.Next()
	}
}
