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
		var tokenString string
		var err error

		// 1. Coba ambil token dari Cookie dulu (Sistem utama)
		tokenString, err = c.Cookie("access_token")

		// 2. JIKA Cookie kosong/gagal (karena diblokir localhost), coba ambil dari Header Authorization (Sistem Cadangan)
		if err != nil || tokenString == "" {
			authHeader := c.GetHeader("Authorization")
			if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
				tokenString = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}

		// Jika kedua cara di atas tetap tidak menemukan token
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token tidak ditemukan atau sesi berakhir!", 
				"status":  401,
			})
			c.Abort()
			return
		}

		// 3. Validasi token ke database
		var accessToken models.PersonalAccessToken
		if err := config.DB.Where("token = ?", tokenString).First(&accessToken).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak valid!", "status": 401})
			c.Abort()
			return
		}

		// 4. Cari data user
		var user models.User
		if err := config.DB.Where("id = ?", accessToken.TokenableID).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "User tidak ditemukan!", "status": 401})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}