package controllers

import (
	"backend-api/config"
	"backend-api/models"
	siswacontroller "backend-api/validations/siswaController"

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

type UpdateSiswaInput struct {
	Nama         string `json:"nama"`
	JenisKelamin string `json:"jenis_kelamin"`
	NoHp         string `json:"no_hp"`
	Alamat       string `json:"alamat"`
	ImageProfile string `json:"image_profile"`
	KelasID      int    `json:"kelas_id"`
	Email        string `json:"email"`
}

func UpdateSiswa(c *gin.Context) {
	var siswa UpdateSiswaInput
	var input siswacontroller.UpdateSiswaValidation
	nis := c.Param("nis")

	if err := config.DB.Table("siswa").Where("nis = ?", nis).First(&siswa).Error; err != nil {
		c.JSON(404, gin.H{"message": "Data siswa dengan nis " + nis + " tidak ditemukan.", "status": 404})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		msg := siswacontroller.TranslateUpdateSiswaError(err)
		c.JSON(400, gin.H{"message": msg, "status": 401})
		return
	}

	config.DB.Table("siswa").Where("nis = ?", nis).Updates(map[string]interface{}{
		"nama":          input.Nama,
		"jenis_kelamin": input.JenisKelamin,
		"no_hp":         input.NoHp,
		"alamat":        input.Alamat,
		"image_profile": input.ImageProfile,
		"kelas_id":      input.KelasId,
	})
	config.DB.Table("users").Where("username = ?", nis).Updates(map[string]interface{}{
		"name":  input.Nama,
		"email": input.Email,
	})
	c.JSON(200, gin.H{
		"success": true,
		"message": "Data siswa berhasil diupdate!",
		"status":  200})
}
