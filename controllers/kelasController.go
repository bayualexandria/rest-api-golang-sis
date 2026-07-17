package controllers

import (
	"backend-api/config"
	"backend-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetKelas(c *gin.Context) {
	kelas := []models.Kelas{}

	if err := config.DB.Model(&kelas).Find(&kelas).Error; err != nil {
		c.JSON(404, gin.H{"message": "Data Kelas Tidak Ada!", "status": "404"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data kelas berhasil ditampilkan!",
		"status":  200,
		"data":    kelas,
		"total":   len(kelas),
	})
}
