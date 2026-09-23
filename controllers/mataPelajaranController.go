package controllers

import (
	"backend-api/config"
	"backend-api/models"

	"github.com/gin-gonic/gin"
)

func GetDataAllMapel(c *gin.Context) {
	var mapel []models.Mapel

	if err := config.DB.Model(&mapel).Find(&mapel).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal mengambil data mata pelajaran",
			"status":  500,
		})
		return
	}

	c.JSON(200, gin.H{
		"data":    mapel,
		"total":   len(mapel),
		"success": true,
		"message": "Data Mata Pelajaran berhasil ditampilkan",
		"status":  200,
	})
}
