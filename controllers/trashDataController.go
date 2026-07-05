package controllers

import (
	"backend-api/config"
	"backend-api/models"

	"github.com/gin-gonic/gin"
)

type UserSiswaAllTrash struct {
	NIS            string `json:"nis"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	JenisKelamin   string `json:"jenis_kelamin"`
	NoHp           string `json:"no_hp"`
	Alamat         string `json:"alamat"`
	ImageProfile   string `json:"image_profile"`
	StatusUserName string `json:"status_user_name"`
}

func GetTrashSiswa(c *gin.Context) {
	var result []UserSiswaAllTrash

	err := config.DB.
		Table("users").
		Joins("JOIN siswa ON users.username = siswa.nis").
		Joins("JOIN status_user ON users.status_id = status_user.id").
		Select(`
		users.username AS nis,
			users.name,
			users.email,
			siswa.jenis_kelamin,
			siswa.no_hp,
			siswa.alamat,
			siswa.nama,
			siswa.image_profile,
			status_user.nama_status AS status_user_name
		`).Where("users.deleted_at IS NOT NULL").
		Scan(&result).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal mengambil trash data siswa",
			"status":  500,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Data trash siswa berhasil ditampilkan!",
		"data":    result,
		"total":   len(result),
		"status":  200,
	})
}

func RestoreDataTrashAllSiswa(c *gin.Context) {
	siswa := config.DB.Unscoped().Model(&models.Siswa{}).Where("deleted_at IS NOT NULL").Update("deleted_at", nil).Error
	user := config.DB.Unscoped().Model(&models.User{}).Where("status_id = ?", 4).Where("deleted_at IS NOT NULL").Update("deleted_at", nil).Error

	if user != nil {
		c.JSON(500, gin.H{
			"message": "Gagal merestore data trash siswa",
			"status":  500,
		})
		return
	}
	if siswa != nil {
		c.JSON(500, gin.H{
			"message": "Gagal merestore data trash siswa",
			"status":  500,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Data trash siswa berhasil direstore!",
		"status":  200,
	})
}

func RestoreDataTrashSiswa(c *gin.Context) {
	nis := c.Param("nis")
	siswa := config.DB.Unscoped().Model(&models.Siswa{}).Where("nis = ?", nis).Where("deleted_at IS NOT NULL").Update("deleted_at", nil).Error
	user := config.DB.Unscoped().Model(&models.User{}).Where("username = ?", nis).Where("status_id = ?", 4).Where("deleted_at IS NOT NULL").Update("deleted_at", nil).Error
	if siswa != nil || user != nil {
		c.JSON(500, gin.H{
			"message": "Gagal merestore data trash siswa",
			"status":  500,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Data trash siswa berhasil direstore!",
		"status":  200,
	})
}

type UserGuruAllTrash struct {
	Nip            string `json:"nip"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	JenisKelamin   string `json:"jenis_kelamin"`
	NoHp           string `json:"no_hp"`
	Alamat         string `json:"alamat"`
	ImageProfile   string `json:"image_profile"`
	StatusUserName string `json:"status_user_name"`
}

func GetTrashGuru(c *gin.Context) {

	var result []UserGuruAllTrash

	err := config.DB.
		Table("users").
		Joins("JOIN guru ON users.username = guru.nip").
		Joins("JOIN status_user ON users.status_id = status_user.id").
		Select(`
		users.username AS nip,
			users.name,
			users.email,
			guru.jenis_kelamin,
			guru.no_hp,
			guru.alamat,
			guru.nama,
			guru.image_profile,
			status_user.nama_status AS status_user_name
		`).Where("users.deleted_at IS NOT NULL").
		Scan(&result).Error

		if err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal mengambil trash data guru",
			"status":  500,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Data trash guru berhasil ditampilkan!",
		"data":    result,
		"total":   len(result),
		"status":  200,
	})

}
