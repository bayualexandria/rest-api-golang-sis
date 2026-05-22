package controllers

import (
	"backend-api/config"
	"backend-api/models"
	guruController "backend-api/validations/guruController"
	"fmt"
	"net/http"
	"os"
	"time"

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
	nip := c.Param("username")
	var guru models.Guru
	if err := config.DB.Where("nip = ?", nip).First(&guru).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Detail Data Guru dengan NIP : " + c.Param("username"),
		"data":    guru,
	})
}

// update guru by nip
func UpdateGuru(c *gin.Context) {
	var guru models.Guru
	var user models.User
	var input guruController.UpdateGuruValidation
	nip := c.Param("nip")

	if err := config.DB.Where("nip", nip).First(&guru).Error; err != nil {
		c.JSON(404, gin.H{
			"message": "Data guru dengan NIP " + nip + " tidak ditemukan",
			"status":  404,
		})
		return
	}

	if err := c.ShouldBind(&input); err != nil {
		msg := guruController.TranslateUpdateGuruError(err)
		c.JSON(400, gin.H{
			"message": "Anda belum merubah data!",
			"data":    msg,
			"status":  400,
		})
		return
	}

	if input.Nama != "" {
		guru.Nama = input.Nama
	}
	if input.JenisKelamin != "" {
		guru.JenisKelamin = input.JenisKelamin
	}
	if input.NoHp != "" {
		guru.NoHp = input.NoHp
	}
	if input.Alamat != "" {
		guru.Alamat = input.Alamat
	}

	if input.ImageProfile != nil {
		file := input.ImageProfile

		os.MkdirAll("storage/guru"+nip, os.ModePerm)
		// Jika gambarnya logo-pendidikan.png

		fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		filePath := "storage/guru/" + nip + "/" + fileName

		// Menghapus file lama
		os.Remove("storage/guru" + nip)
		os.Remove(guru.ImageProfile)

		c.SaveUploadedFile(file, filePath)
		guru.ImageProfile = filePath

	}

	if err := config.DB.Save(&guru).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal mengupdate database: " + err.Error()})
		return
	}

	config.DB.Model(&user).Where("username", nip).Updates(map[string]interface{}{
		"name": guru.Nama,
	})
	c.JSON(200, gin.H{
		"success": true,
		"message": "Data siswa berhasil diupdate",
		"status":  200,
	})

}
