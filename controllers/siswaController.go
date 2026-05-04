package controllers

import (
	"backend-api/config"
	"backend-api/models"
	siswacontroller "backend-api/validations/siswaController"

	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type UserAll struct {
	NIS            string `json:"nis"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	JenisKelamin   string `json:"jenis_kelamin"`
	NoHp           string `json:"no_hp"`
	Alamat         string `json:"alamat"`
	ImageProfile   string `json:"image_profile"`
	StatusUserName string `json:"status_user_name"`
}

func GetSiswa(c *gin.Context) {
	var result []UserAll

	err := config.DB.
		Table("users").
		Joins("JOIN siswa ON users.username = siswa.nis").
		Joins("JOIN status_user ON users.status_user_id = status_user.id").
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
		`).
		Scan(&result).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal mengambil data siswa",
			"status":  500,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Data siswa berhasil ditampilkan!",
		"data":    result,
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

func UpdateSiswa(c *gin.Context) {
	var input siswacontroller.UpdateSiswaValidation
	var siswa models.Siswa
	var user models.User

	nis := c.Param("nis")

	// cek data siswa
	if err := config.DB.Where("nis = ?", nis).First(&siswa).Error; err != nil {
		c.JSON(404, gin.H{
			"message": "Data siswa dengan NIS " + nis + " tidak ditemukan",
			"status":  404,
		})
		return
	}
	// bind form-data
	if err := c.ShouldBind(&input); err != nil {
		msg := siswacontroller.TranslateUpdateSiswaError(err)
		c.JSON(400, gin.H{
			"message": msg,
			"status":  400,
		})
		return
	}

	// update field biasa
	siswa.Nama = input.Nama
	siswa.JenisKelamin = input.JenisKelamin
	siswa.NoHp = input.NoHp
	siswa.Alamat = input.Alamat

	// handle upload gambar (optional)
	if input.ImageProfile != nil {
		file := input.ImageProfile
		// Jika folder storages belum ada, buat folder tersebut
		if err := os.MkdirAll("storages/siswa/"+nis, os.ModePerm); err != nil {
			c.JSON(500, gin.H{
				"message": "Gagal membuat folder penyimpanan",
				"status":  500,
			})
			return
		}

		// buat nama file unik
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		filePath := "storages/siswa/" + nis + "/" + filename

		// simpan file
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(500, gin.H{
				"message": "Gagal upload gambar",
				"status":  500,
			})
			return
		}

		// simpan path ke database
		siswa.ImageProfile = filePath
	}

	// simpan ke DB
	if err := config.DB.Model(&siswa).Updates(map[string]interface{}{
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
	if err := config.DB.Model(&user).Where("username", nis).Updates(map[string]interface{}{
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
