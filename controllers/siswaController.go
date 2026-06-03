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
	nis := c.Param("username")

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

func AddSiswa(c *gin.Context) {
	var input siswacontroller.AddSiswaValidation

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

	// Buat instance baru dari model Siswa

	siswa := models.Siswa{
		Nis:          input.NIS,
		Nama:         input.Nama,
		JenisKelamin: input.JenisKelamin,
		NoHp:         input.NoHp,
		Alamat:       input.Alamat,
	}

	user := models.User{
		Username: fmt.Sprintf("%d", input.NIS),
		Name:     input.Nama,
		Email:    input.Email,
		StatusId: 4, // ID untuk status "siswa"
	}

	// Simpan ke database
	if err := config.DB.Create(&siswa).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal menambahkan data siswa!",
			"status":  500,
		})
		return
	}
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal menambahkan data user untuk siswa!",
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
		os.Remove("storage/guru" + nis)
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
	config.DB.Model(&user).Where("username", nis).Updates(map[string]interface{}{
		"name": siswa.Nama,
	})

	c.JSON(200, gin.H{
		"success": true,
		"message": "Data siswa berhasil diupdate",
		"status":  200,
	})
}
