package controllers

import (
	"backend-api/config"
	"backend-api/models"
	"backend-api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginUserSocialMedia(c *gin.Context) {
	var user models.User
	email := c.Param("email")
	idGoogle := c.Param("idGoogle")
	nameGoogle := c.Param("nameGoogle")
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(403, gin.H{"message": "User belum terdaftar! Silahkan hubungi administrator sekolah.", "status": 403})
		return
	} else {

		if err := config.DB.Where("status_id != ?", 4).First(&user).Error; err != nil {
			c.JSON(403, gin.H{"message": "User ini bukan siswa! Silahkan hubungi administrator sekolah.", "status": 403})
			return
		} else {

			if user.EmailVerifiedAt == "" {
				c.JSON(403, gin.H{"message": "Email belum terverifikasi! Silahkan verifikasi email terlebih dahulu.", "status": 403})
				return
			} else {
				var linkedSocialAccount models.LinkedSocialMediaAccount
				token, err := utils.GenerateJWT(user.Username)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memuat token"})
					return
				}
				var inputToken models.PersonalAccessToken
				inputToken.Token = token
				inputToken.TokenableType = "User"
				inputToken.TokenableID = user.Username
				inputToken.Name = "Personal Access Token"
				inputToken.Abilities = "*"
				inputToken.LastUsedAt = time.Now().Format("2006-01-02 15:04:05")
				inputToken.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
				inputToken.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

				if err := config.DB.Where("provider_id = ?", idGoogle).First(&linkedSocialAccount).Error; err != nil {
					linkedSocialAccount.UserID = user.Username
					linkedSocialAccount.ProviderName = nameGoogle
					linkedSocialAccount.ProviderID = idGoogle
					config.DB.Create(&linkedSocialAccount)
					config.DB.Create(&inputToken)
					// Di dalam controller saat LOGIN BERHASIL:
					// Gunakan properti cookie standar go untuk mengatur SameSite secara eksplisit
					c.Writer.Header().Set("Set-Cookie", "access_token="+token+"; Max-Age=86400; Path=/; SameSite=Lax; HttpOnly")

					// ATAU jika menggunakan c.SetCookie bawaan Gin, pastikan parameternya seperti ini:
					// c.SetCookie("access_token", token, 86400, "/", "", false, true)
					c.JSON(200, gin.H{"message": "Login berhasil!", "status": 200, "user": user})
					return
				} else {
					linkedSocialAccount.ProviderName = nameGoogle
					linkedSocialAccount.ProviderID = idGoogle
					config.DB.Save(&linkedSocialAccount)
					config.DB.Create(&inputToken)
					// Di dalam controller saat LOGIN BERHASIL:
					// Gunakan properti cookie standar go untuk mengatur SameSite secara eksplisit
					c.Writer.Header().Set("Set-Cookie", "access_token="+token+"; Max-Age=86400; Path=/; SameSite=Lax; HttpOnly")

					// ATAU jika menggunakan c.SetCookie bawaan Gin, pastikan parameternya seperti ini:
					// c.SetCookie("access_token", token, 86400, "/", "", false, true)
					c.JSON(200, gin.H{"message": "Login berhasil!", "status": 200, "user": user})
					return
				}
			}
		}

	}

}

// Implementation for social media login goes here
