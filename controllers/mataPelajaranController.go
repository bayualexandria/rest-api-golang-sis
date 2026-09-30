package controllers

import (
	"backend-api/config"
	"backend-api/models"
	"backend-api/validations/mapelController"

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

func InsertDataMapel(c *gin.Context) {
	var input mapelController.AddMapelValidation
	var mapel models.Mapel

	// bind form-data
	if err := c.ShouldBind(&input); err != nil {
		msg := mapelController.TranslateAddMapelError(err)
		c.JSON(400, gin.H{
			"message": "Gagal menambahkan data mapel!",
			"data":    msg,
			"status":  400,
		})
		return
	}

	if err := config.DB.Model(&mapel).Where("nama = ?", input.Nama).First(&mapel).Error; err == nil {
		c.JSON(401, gin.H{
			"message": "Nama mapel sudah digunakan!",
			"status":  401,
		})
		return
	}

	if err := config.DB.Model(&mapel).Where("kode = ?", input.Kode).First(&mapel).Error; err == nil {
		c.JSON(401, gin.H{
			"message": "Kode mapel sudah digunakan!",
			"status":  401,
		})
		return
	}

	if err := config.DB.Model(&mapel).Create(map[string]interface{}{
		"nama":      input.Nama,
		"kode":      input.Kode,
		"deskripsi": input.Deskripsi,
	}).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal menambahkan data mapel!",
			"status":  500,
		})
		return
	}

	c.JSON(201, gin.H{
		"success": true,
		"message": "Data mapel berhasil ditambahkan!",
		"data": gin.H{
			"nama":      input.Nama,
			"kode":      input.Kode,
			"deskripsi": input.Deskripsi,
		},
	})

}
