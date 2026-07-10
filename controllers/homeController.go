package controllers

import (
	"backend-api/config"
	"backend-api/models"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"data":    []string{},
		"message": "Data berhasil ditampilkan!",
		"success": true,
	})
}

func ProfileSekolahHandler(c *gin.Context) {
	var sekolah models.ProfileSekolah
	err := config.DB.First(&sekolah).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal mengambil data profile sekolah",
			"status":  500,
		})
		return
	}
	c.JSON(200, gin.H{
		"data":    sekolah,
		"message": "Data berhasil ditampilkan!",
		"success": true,
	})
}
