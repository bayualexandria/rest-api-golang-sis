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
	nis := c.Param("nis")

	var result UserWithSiswa
	// Join dengan tabel siswa berdasarkan nis
	config.DB.Table("users").Joins("JOIN siswa ON users.username = siswa.nis").Joins("JOIN status_user ON users.status_user_id = status_user.id").Where("users.username = ?", nis).Select(" users.name, users.email,users.username AS nis,  siswa.jenis_kelamin, siswa.no_hp, siswa.alamat, siswa.image_profile, status_user.nama_status AS status_user_name").First(&result)

	c.JSON(http.StatusOK, gin.H{"data": result})
}
