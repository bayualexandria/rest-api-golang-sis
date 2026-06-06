package controllers

import (
	"backend-api/config"
	"backend-api/models"
	"backend-api/utils"
	guruController "backend-api/validations/guruController"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type UserAllGuru struct {
	Nip            int    `json:"nip"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	JenisKelamin   string `json:"jenis_kelamin"`
	NoHp           string `json:"no_hp"`
	Alamat         string `json:"alamat"`
	ImageProfile   string `json:"image_profile"`
	StatusUserName string `json:"status_user_name"`
}

func GetGuru(c *gin.Context) {
	var guru []UserAllGuru

	err := config.DB.
		Table("users").
		Joins("JOIN guru ON users.username = guru.nip").
		Joins("JOIN status_user ON users.status_id = status_user.id").
		Select(`
		guru.nip,
			users.name,
			users.email,
			guru.jenis_kelamin,
			guru.no_hp,
			guru.alamat,
			guru.nama,
			guru.image_profile,
			status_user.nama_status AS status_user_name
		`).Where("users.deleted_at IS NULL").
		Scan(&guru).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal mengambil data guru",
			"status":  500,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Data guru berhasil ditampilkan!",
		"data":    guru,
		"total":   len(guru),
	})
}

func AddGuru(c *gin.Context) {
	var input guruController.AddGuruValidation
	var guru models.Guru
	var user models.User

	// bind form-data
	if err := c.ShouldBind(&input); err != nil {
		msg := guruController.TranslateAddGuruError(err)
		c.JSON(400, gin.H{
			"message": "Gagal menambahkan data guru!",
			"data":    msg,
			"status":  400,
		})
		return
	}

	if input.ImageProfile != nil {
		file := input.ImageProfile
		// Jika folder storages belum ada, buat folder tersebut
		os.MkdirAll("storage/guru/"+fmt.Sprintf("%d", input.Nip), os.ModePerm)

		// buat nama file unik
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		filePath := "storage/guru/" + fmt.Sprintf("%d", input.Nip) + "/" + filename

		// simpan file
		os.Remove("storage/guru" + fmt.Sprintf("%d", input.Nip))
		os.Remove(guru.ImageProfile)

		c.SaveUploadedFile(file, filePath)

		// simpan path ke database
		guru.ImageProfile = filePath
	} else {

		guru.ImageProfile = "/storage/logo-pendidikan.png"
	}
	hashPassword, err := utils.HashPassword(fmt.Sprintf("%d", input.Nip))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal menghash password!",
			"status":  500,
		})
		return
	}
	// Simpan ke database
	if err := config.DB.Model(&user).Create(map[string]interface{}{
		"username":          input.Nip,
		"name":              input.Nama,
		"email":             input.Email,
		"email_verified_at": time.Now().Format("2006-01-02 15:04:05"),
		"password":          hashPassword,     // Ganti dengan password default atau generate secara acak
		"status_id":         input.StatusId, // Misalnya 4 adalah ID untuk status "siswa"
	}).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Email atau Username sudah digunakan!",
			"status":  500,
		})
		return
	}
	if err := config.DB.Model(&guru).Create(map[string]interface{}{
		"nip":           input.Nip,
		"nama":          input.Nama,
		"jenis_kelamin": input.JenisKelamin,
		"no_hp":         input.NoHp,
		"alamat":        input.Alamat,
		"image_profile": guru.ImageProfile,
	}).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal menambahkan data guru!",
			"status":  500,
		})
		return
	}

	if input.StatusId == "1" {
		input.StatusId = "Admin"
	} else if input.StatusId == "2" {
		input.StatusId = "Wali Kelas"
	} else if input.StatusId == "3" {
		input.StatusId = "Guru"
	}

	c.JSON(201, gin.H{
		"success": true,
		"message": "Data guru berhasil ditambahkan!",
		"data": gin.H{
			"nip":           input.Nip,
			"nama":          input.Nama,
			"jenis_kelamin": input.JenisKelamin,
			"no_hp":         input.NoHp,
			"alamat":        input.Alamat,
			"image_profile": guru.ImageProfile,
			"email":         input.Email,
			"status_user":   input.StatusId,
		},
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
		"message": "Data guru berhasil diupdate",
		"status":  200,
	})

}

func DeleteGuru(c *gin.Context) {
	nip := c.Param("nip")
	var guru models.Guru
	var user models.User
	// cek data guru
	if err := config.DB.Where("nip = ?", nip).First(&guru).Error; err != nil {
		c.JSON(404, gin.H{
			"message": "Data guru dengan NIP " + nip + " tidak ditemukan",
			"status":  404,
		})
		return
	}
	// hapus data guru
	if err := config.DB.Model(&guru).Where("nip = ?", nip).Delete(&guru).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal menghapus data guru!",
			"status":  500,
		})
		return
	}

	// hapus data user terkait
	if err := config.DB.Model(&user).Where("username", nip).Delete(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal menghapus data user terkait!",
			"status":  500,
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"message": "Data guru berhasil dihapus!",
		"status":  200,
	})
}
