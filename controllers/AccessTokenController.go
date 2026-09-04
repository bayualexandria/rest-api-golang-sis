package controllers

import (
	"backend-api/config"
	"backend-api/models"

	"github.com/gin-gonic/gin"
)

func GetAccessToken(c *gin.Context) {
	personalAccessToken := models.PersonalAccessToken{}
	// Implementation for getting access token
	username := c.Param("username")

	if err := config.DB.Table("personal_access_tokens").Where("tokenable_id = ?", username).Find(&personalAccessToken).Error; err != nil {
		c.JSON(404, gin.H{"message": "Access token tidak ditemukan!", "status": 404})
		return
	}

	// Hapus data access token bedasarkan username jika sudah ada
	config.DB.Table("personal_access_tokens").Where("tokenable_id = ?", username).Delete(&models.PersonalAccessToken{})

	c.JSON(200, personalAccessToken)

}
