package controllers

import (
	"backend-api/config"
	"backend-api/models"
	"backend-api/utils"

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
		var siswa models.Siswa
		if err := config.DB.Where("nis = ?", user.Username).First(&siswa).Error; err != nil {
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
					c.JSON(500, gin.H{"message": "Error generating token", "status": 500})
					return
				}
				if err := config.DB.Where("provider_id = ?", idGoogle).First(&linkedSocialAccount).Error; err != nil {
					linkedSocialAccount.UserID = uint(user.ID)
					linkedSocialAccount.ProviderName = nameGoogle
					linkedSocialAccount.ProviderID = idGoogle
					config.DB.Create(&linkedSocialAccount)
					c.JSON(200, gin.H{"message": "Login berhasil!", "status": 200, "token": token, "data": user})
				} else {
					linkedSocialAccount.ProviderName = nameGoogle
					linkedSocialAccount.ProviderID = idGoogle
					config.DB.Save(&linkedSocialAccount)
					c.JSON(200, gin.H{"message": "Login berhasil!", "status": 200, "token": token, "data": user})
				}
			}
		}

	}

}

// Implementation for social media login goes here
