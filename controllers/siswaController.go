package controllers

import (
	"backend-api/config"
	"backend-api/models"
	"backend-api/utils"
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

func AddSiswa(c *gin.Context) {
	var input siswacontroller.AddSiswaValidation
	var siswa models.Siswa
	var user models.User

	// bind form-data
	if err := c.ShouldBind(&input); err != nil {
		msg := siswacontroller.TranslateAddSiswaError(err)
		c.JSON(400, gin.H{
			"message": "Gagal menambahkan data siswa!",
			"data":    msg,
			"status":  400,
		})
		return
	}

	if input.ImageProfile != nil {
		file := input.ImageProfile
		// Jika folder storages belum ada, buat folder tersebut
		os.MkdirAll("storage/siswa/"+fmt.Sprintf("%d", input.Nis), os.ModePerm)

		// buat nama file unik
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		filePath := "storage/siswa/" + fmt.Sprintf("%d", input.Nis) + "/" + filename

		c.SaveUploadedFile(file, filePath)

		// simpan path ke database
		siswa.ImageProfile = filePath
	}
	siswa.ImageProfile = "/storage/logo-pendidikan.png"
	hashPassword, err := utils.HashPassword(fmt.Sprintf("%d", input.Nis))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal menghash password!",
			"status":  500,
		})
		return
	}
	// Simpan ke database
	if err := config.DB.Model(&user).Create(map[string]interface{}{
		"username":  input.Nis,
		"name":      input.Nama,
		"email":     input.Email,
		"password":  hashPassword, // Ganti dengan password default atau generate secara acak
		"status_id": 4,            // Misalnya 4 adalah ID untuk status "siswa"
	}).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Email atau Username sudah digunakan!",
			"status":  500,
		})
		return
	}
	if err := config.DB.Model(&siswa).Create(map[string]interface{}{
		"nis":           input.Nis,
		"nama":          input.Nama,
		"jenis_kelamin": input.JenisKelamin,
		"no_hp":         input.NoHp,
		"alamat":        input.Alamat,
		"image_profile": siswa.ImageProfile,
	}).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal menambahkan data siswa!",
			"status":  500,
		})
		return
	}

	c.JSON(201, gin.H{
		"success": true,
		"message": "Data siswa berhasil ditambahkan!",
		"data": gin.H{
			"nis":           siswa.Nis,
			"nama":          siswa.Nama,
			"jenis_kelamin": siswa.JenisKelamin,
			"no_hp":         siswa.NoHp,
			"alamat":        siswa.Alamat,
			"image_profile": siswa.ImageProfile,
			"email":         user.Email,
			"status_user":   "siswa",
		},
	})
}

func UpdateSiswa(c *gin.Context) {
	var siswa models.Siswa
	var input siswacontroller.UpdateSiswaValidation
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
			"message": "Anda belum merubah data!",
			"data":    msg,
			"status":  400,
		})
		return
	}

	// update field biasa
	if input.Nama != "" {
		siswa.Nama = input.Nama
	}
	if input.JenisKelamin != "" {
		siswa.JenisKelamin = input.JenisKelamin
	}
	if input.NoHp != "" {
		siswa.NoHp = input.NoHp
	}
	if input.Alamat != "" {
		siswa.Alamat = input.Alamat
	}

	// handle upload gambar (optional)
	if input.ImageProfile != nil {
		file := input.ImageProfile
		// Jika folder storages belum ada, buat folder tersebut
		os.MkdirAll("storage/siswa/"+nis, os.ModePerm)

		// buat nama file unik
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		filePath := "storage/siswa/" + nis + "/" + filename

		// simpan file
		os.Remove("storage/siswa" + nis)
		os.Remove(siswa.ImageProfile)

		c.SaveUploadedFile(file, filePath)

		// simpan path ke database
		siswa.ImageProfile = filePath
	}

	// simpan ke DB
	if err := config.DB.Save(&siswa).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal mengupdate database: " + err.Error()})
		return
	}
	if err := config.DB.Model(&user).Where("username", nis).Updates(map[string]interface{}{
		"name": siswa.Nama,
	}).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal mengupdate data user: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Data siswa berhasil diupdate",
		"status":  200,
	})
}

func DeleteSiswa(c *gin.Context) {
	nis := c.Param("nis")
	var siswa models.Siswa
	// cek data siswa
	if err := config.DB.Where("nis = ?", nis).First(&siswa).Error; err != nil {
		c.JSON(404, gin.H{
			"message": "Data siswa dengan NIS " + nis + " tidak ditemukan",
			"status":  404,
		})
		return
	}
	// hapus data siswa
	if err := config.DB.Delete(&siswa).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal menghapus data siswa!",
			"status":  500,
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"message": "Data siswa berhasil dihapus!",
		"status":  200,
	})
}
