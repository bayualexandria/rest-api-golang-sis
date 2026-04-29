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

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak ditemukan!", "status": 401})
			c.Abort()
			return
		}

		if !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Format token salah!", "status": 401})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(token, "Bearer ")

		var accessToken models.PersonalAccessToken
		if err := config.DB.Where("token = ?", tokenString).First(&accessToken).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak valid!", "status": 401})
			c.Abort()
			return
		}

		var user models.User
		if err := config.DB.Where("id = ?", accessToken.TokenableID).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "User tidak ditemukan!", "status": 401})
			c.Abort()
			return
		}

		// cek status user
		if user.StatusUserId != "4" {
			c.JSON(http.StatusForbidden, gin.H{"message": "User ini tidak memiliki akses!", "status": 403})
			c.Abort()
			return
		}

		// simpan user ke context (optional tapi penting)
		c.Set("user", user)

		c.Next()
	}
}

