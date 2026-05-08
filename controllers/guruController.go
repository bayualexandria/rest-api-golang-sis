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
	nip := c.Param("nip")
	var guru models.Guru
	if err := config.DB.Where("nip = ?", nip).First(&guru).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Detail Data Guru dengan NIP : " + c.Param("nip"),
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
		msg := guruController.TranslateUpdateSiswaError(err)
		c.JSON(400, gin.H{
			"message": msg,
			"status":  400,
		})
		return
	}

	guru.Nama = input.Nama
	guru.JenisKelamin = input.JenisKelamin
	guru.NoHp = input.NoHp
	guru.Alamat = input.Alamat

	if input.ImageProfile != nil {
		file := input.ImageProfile

		if err := os.MkdirAll("storages/guru"+nip, os.ModePerm); err != nil {
			c.JSON(500, gin.H{
				"message": "Gagal membuat folder penyimpanan",
				"status":  500,
			})
			return
		}

		fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		filePath := "storages/guru/" + nip + "/" + fileName

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(500, gin.H{
				"message": "Gagal upload gambar",
				"status":  500,
			})
			return
		}
		guru.ImageProfile = filePath

	}
	if err := config.DB.Model(&guru).Updates(map[string]interface{}{
		"nama":          input.Nama,
		"jenis_kelamin": input.JenisKelamin,
		"alamat":        input.Alamat,
		"no_hp":         input.NoHp,
	}).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal mengupdate data siswa",
			"status":  500,
		})
		return
	}

	if err := config.DB.Model(&user).Where("username", nip).Updates(map[string]interface{}{
		"name":  input.Nama,
		"email": input.Email,
	}).Error; err != nil {
		c.JSON(401, gin.H{"message": "Email yang anda masukan sudah terdaftar!", "status": 401})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"message": "Data siswa berhasil diupdate",
		"status":  200,
	})

}
