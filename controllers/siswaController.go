package controllers

import (
	"backend-api/config"
	"backend-api/models"

	"github.com/gin-gonic/gin"
)

func GetSiswa(c *gin.Context) {
	var siswa []models.Siswa

	// Ambil semua data dari database
	result := config.DB.Find(&siswa)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Terjadi kesalahan saat mengambil data siswa.",
			"error":   result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Data siswa tidak ditemukan.",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Data siswa berhasil ditampilkan!",
		"data":    siswa,
	})
}

func GetSiswaByID(c *gin.Context) {
	var siswa models.Siswa
	nis := c.Param("nis")

	// Ambil data berdasarkan ID dari database
	result := config.DB.Where("nis = ?", nis).First(&siswa)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": "Siswa dengan nis " + nis + " tidak ditemukan.",
			"status":  404,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Data siswa berhasil ditampilkan!",
		"data":    siswa,
	})
}
