package controllers

import (
	"backend-api/config"
	"backend-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

type UserWithSiswa struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	NIS            string `json:"nis"`
	JenisKelamin   string `json:"jenis_kelamin"`
	NoHp           string `json:"no_hp"`
	Alamat         string `json:"alamat"`
	ImageProfile   string `json:"image_profile"`
	StatusUserName string `json:"status_user_name"`
}

func GetUsersByNIS(c *gin.Context) {
	nis := c.Param("username")

	var result UserWithSiswa
	// Join dengan tabel siswa berdasarkan nis
	siswa := config.DB.Table("users").Joins("JOIN siswa ON users.username = siswa.nis").Joins("JOIN status_user ON users.status_id = status_user.id").Where("users.username = ?", nis).Where("users.deleted_at IS NULL").Select(" users.name, users.email,users.username AS nis,  siswa.jenis_kelamin, siswa.no_hp, siswa.alamat, siswa.image_profile, status_user.nama_status AS status_user_name").First(&result)
	if siswa.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Siswa tidak ditemukan atau NIS salah", "status": 404})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

type UserWithGuru struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	NIP            string `json:"nip" gorm:"column:nip"`
	JenisKelamin   string `json:"jenis_kelamin"`
	NoHp           string `json:"no_hp"`
	Alamat         string `json:"alamat"`
	ImageProfile   string `json:"image_profile"`
	StatusUserName string `json:"status_user_name"`
	StatusID       int    `json:"status_id"`
}

func GetUsersByNIP(c *gin.Context) {
	nip := c.Param("username")

	var result UserWithGuru
	// Join dengan tabel guru berdasarkan nip
	err := config.DB.Table("users").
		Joins("JOIN guru ON users.username = guru.nip").
		Joins("JOIN status_user ON users.status_id = status_user.id").
		Where("users.username = ?", nip).
		Where("users.deleted_at IS NULL").
		Select("users.name, users.email, guru.nip AS nip, guru.jenis_kelamin, guru.no_hp, guru.alamat, guru.image_profile, status_user.nama_status AS status_user_name, users.status_id").
		First(&result).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Guru tidak ditemukan atau NIP salah", "status": 404})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetUsersByUsername(c *gin.Context) {
	username := c.Param("username")

	var user models.User
	config.DB.Table("users").
		Where("username = ?", username).
		First(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}
