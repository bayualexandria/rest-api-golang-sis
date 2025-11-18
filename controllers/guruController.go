package controllers

import (
	"backend-api/config"
	"backend-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGuru(c *gin.Context) {
	var guru []models.Guru

	result := config.DB.Find(&guru)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Terjadi kesalahan saat mengambil data guru.",
			"error":   result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Data guru tidak ditemukan.",
		})
		return
	}

	c.JSON(200, gin.H{
		"succes":  true,
		"message": "Data guru berhasil ditampilkan!",
		"data":    guru,
	})
}

// get post by id
func GetGuruById(c *gin.Context) {
	var guru models.Guru
	if err := config.DB.Where("nip = ?", c.Param("nip")).First(&guru).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Detail Data Guru dengan NIP : " + c.Param("nip"),
		"data":    guru,
	})
}
