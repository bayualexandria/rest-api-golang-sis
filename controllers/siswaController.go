package controllers

import (
	"backend-api/config"
	"backend-api/models"
	"backend-api/notifications"
	"backend-api/utils"
	siswacontroller "backend-api/validations/siswaController"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type UserAllSiswa struct {
	ID             int    `json:"id"`
	NIS            string `json:"nis"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Ttl            string `json:"ttl"`
	JenisKelamin   string `json:"jenis_kelamin"`
	NoHp           string `json:"no_hp"`
	Alamat         string `json:"alamat"`
	ImageProfile   string `json:"image_profile"`
	StatusUserName string `json:"status_user_name"`
}

func GetSiswa(c *gin.Context) {
	var result []UserAllSiswa

	err := config.DB.
		Table("users").
		Joins("JOIN siswa ON users.username = siswa.nis").
		Joins("JOIN status_user ON users.status_id = status_user.id").
		Select(`
		siswa.id,
		users.username AS nis,
			users.name,
			users.email,
			siswa.ttl,
			siswa.jenis_kelamin,
			siswa.no_hp,
			siswa.alamat,
			siswa.nama,
			siswa.image_profile,
			status_user.nama_status AS status_user_name
		`).Where("users.deleted_at IS NULL").
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
		"total":   len(result),
	})
}

func GetSiswaSearch(c *gin.Context) {
	var result []UserAllSiswa

	err := config.DB.
		Table("users").
		Joins("JOIN siswa ON users.username = siswa.nis").
		Joins("JOIN status_user ON users.status_id = status_user.id").
		Select(`
		siswa.id,
		users.username AS nis,
			users.name,
			users.email,
			siswa.ttl,
			siswa.jenis_kelamin,
			siswa.no_hp,
			siswa.alamat,
			siswa.nama,
			siswa.image_profile,
			status_user.nama_status AS status_user_name
		`).Where("users.deleted_at IS NULL").Where("siswa.status_siswa_id = ?", 2).
		Scan(&result).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal mengambil data search siswa",
			"status":  500,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Data siswa berhasil ditampilkan!",
		"data":    result,
		"total":   len(result),
	})
}

func AddSiswa(c *gin.Context) {
	var input siswacontroller.AddSiswaValidation
	var siswa models.Siswa
	var user models.User

	// bind form-data
	if err := c.ShouldBindJSON(&input); err != nil {
		msg := siswacontroller.TranslateAddSiswaError(err)
		c.JSON(400, gin.H{
			"message": "Gagal menambahkan data siswa!",
			"data":    msg,
			"status":  400,
		})
		return
	}

	// Simpan ke database
	if err := config.DB.Model(&user).Create(map[string]interface{}{
		"username":          input.Nis,
		"name":              input.Nama,
		"email":             input.Email,
		"email_verified_at": time.Now().Format("2006-01-02 15:04:05"),
		"password":          utils.HashPasswordUser("admin123"), // Ganti dengan password default atau generate secara acak
		"status_id":         "4",                                // Misalnya 4 adalah ID untuk status "siswa"
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
		"ttl":           input.Ttl,
		"jenis_kelamin": input.JenisKelamin,
		"no_hp":         input.NoHp,
		"alamat":        input.Alamat,
		"image_profile": "storage/logo-pendidikan.png",
	}).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal menambahkan data siswa!",
			"status":  500,
		})
		return
	}

	token, err := utils.GenerateJWT(input.Nis)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memuat token"})
		return
	}
	var inputToken models.PersonalAccessToken
	inputToken.Token = token
	inputToken.TokenableType = "User"
	inputToken.TokenableID = input.Nis
	inputToken.Name = "Personal Access Token"
	inputToken.Abilities = "*"

	config.DB.Create(&inputToken)
	notifications.NotifikasiAktivasiAkunUser(input.Email, input.Nama, "Selamat akun anda telah berhasil dibuat. Silahkan verifikasi email anda untuk mengaktifkan akun anda, dengan cara klik link dibawah ini: ", os.Getenv("APP_URL")+"/api/auth/verify/"+input.Email+"/"+token)

	c.JSON(201, gin.H{
		"success": true,
		"message": "Data siswa berhasil ditambahkan!",
		"data": gin.H{
			"nis":           input.Nis,
			"nama":          input.Nama,
			"ttl":           input.Ttl,
			"jenis_kelamin": input.JenisKelamin,
			"no_hp":         input.NoHp,
			"alamat":        input.Alamat,
			"image_profile": "storage/logo-pendidikan.png",
			"email":         input.Email,
			"status_user":   "siswa",
		},
	})
}

func GetSiswaByNIS(c *gin.Context) {
	nis := c.Param("username")

	var result UserWithSiswa
	// Join dengan tabel siswa berdasarkan nis
	siswa := config.DB.Table("users").
		Joins("JOIN siswa ON users.username = siswa.nis").
		Joins("JOIN status_user ON users.status_id = status_user.id").
		Where("users.username = ?", nis).
		Where("users.deleted_at IS NULL").
		Select(" users.name, users.email,users.username AS nis,siswa.ttl AS tempat_tanggal_lahir,  siswa.jenis_kelamin, siswa.no_hp, siswa.alamat, siswa.image_profile, status_user.nama_status AS status_user_name").
		First(&result)
	if siswa.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Siswa tidak ditemukan atau NIS salah", "status": 404})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func UpdateSiswa(c *gin.Context) {
	var siswa models.Siswa
	var user models.User
	var input siswacontroller.UpdateSiswaValidation

	nis := c.Param("nis")

	// =========================================================
	// CARI DATA GURU
	// =========================================================

	if err := config.DB.
		Where("nis = ?", nis).
		First(&siswa).Error; err != nil {

		c.JSON(404, gin.H{
			"message": "Data guru dengan NIP " + nis + " tidak ditemukan",
			"status":  404,
		})

		return
	}

	// =========================================================
	// CARI DATA USER
	// =========================================================

	if err := config.DB.
		Where("username = ?", nis).
		First(&user).Error; err != nil {

		c.JSON(404, gin.H{
			"message": "Data user dengan username " + nis + " tidak ditemukan",
			"status":  404,
		})

		return
	}

	// =========================================================
	// BIND FORM DATA
	// =========================================================

	if err := c.ShouldBind(&input); err != nil {
		msg := siswacontroller.TranslateUpdateSiswaError(err)

		c.JSON(400, gin.H{
			"message": "Anda belum merubah data!",
			"data":    msg,
			"status":  400,
		})

		return
	}

	// =========================================================
	// UPDATE DATA GURU
	// =========================================================

	if input.Nama != "" {
		siswa.Nama = input.Nama
	}

	if input.Ttl != "" {
		siswa.Ttl = input.Ttl
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

	// =========================================================
	// UPDATE EMAIL USER
	// =========================================================

	if input.Email != "" {
		user.Email = input.Email
	}

	// =========================================================
	// UPDATE FOTO PROFILE
	// =========================================================

	if input.ImageProfile != nil {
		file := input.ImageProfile

		// Folder penyimpanan berdasarkan NIP
		folderPath := filepath.Join(
			"storage",
			"guru",
			nis,
		)

		// Buat folder jika belum ada
		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			c.JSON(500, gin.H{
				"message": "Gagal membuat folder penyimpanan gambar",
				"error":   err.Error(),
				"status":  500,
			})

			return
		}

		// =====================================================
		// HAPUS GAMBAR LAMA
		// =====================================================

		oldImagePath := filepath.Clean(siswa.ImageProfile)

		// Logo default TIDAK BOLEH DIHAPUS
		defaultImagePath := filepath.Clean(
			"storage/logo-pendidikan.png",
		)

		if oldImagePath != "" &&
			oldImagePath != "." &&
			oldImagePath != defaultImagePath {

			if _, err := os.Stat(oldImagePath); err == nil {
				if err := os.Remove(oldImagePath); err != nil {
					fmt.Println(
						"Gagal menghapus gambar lama:",
						err,
					)
				}
			}
		}

		// =====================================================
		// BUAT NAMA FILE BARU
		// =====================================================

		fileName := fmt.Sprintf(
			"%d_%s",
			time.Now().UnixNano(),
			file.Filename,
		)

		filePath := filepath.Join(
			folderPath,
			fileName,
		)

		// =====================================================
		// SIMPAN FOTO BARU
		// =====================================================

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(500, gin.H{
				"message": "Gagal menyimpan gambar",
				"error":   err.Error(),
				"status":  500,
			})

			return
		}

		// Simpan path gambar baru
		siswa.ImageProfile = filePath
	}

	// =========================================================
	// UPDATE DATABASE GURU
	// =========================================================

	if err := config.DB.
		Model(&siswa).
		Where("nis = ?", nis).
		Updates(&siswa).Error; err != nil {

		c.JSON(500, gin.H{
			"message": "Gagal mengupdate database guru",
			"error":   err.Error(),
			"status":  500,
		})

		return
	}

	// =========================================================
	// UPDATE DATABASE USER
	// =========================================================

	if err := config.DB.
		Model(&user).
		Where("username = ?", nis).
		Updates(map[string]interface{}{
			"name":  siswa.Nama,
			"email": user.Email,
		}).Error; err != nil {

		c.JSON(500, gin.H{
			"message": "Gagal mengupdate data user",
			"error":   err.Error(),
			"status":  500,
		})

		return
	}

	// =========================================================
	// RESPONSE
	// =========================================================

	c.JSON(200, gin.H{
		"success": true,
		"data":    siswa,
		"message": "Data guru berhasil diupdate",
		"status":  200,
	})
}

func DeleteSiswa(c *gin.Context) {
	nis := c.Param("nis")
	var siswa models.Siswa
	var user models.User
	// cek data siswa
	if err := config.DB.Where("nis = ?", nis).First(&siswa).Error; err != nil {
		c.JSON(404, gin.H{
			"message": "Data siswa dengan NIS " + nis + " tidak ditemukan",
			"status":  404,
		})
		return
	}
	// hapus data siswa
	if err := config.DB.Model(&siswa).Where("nis = ?", nis).Delete(&siswa).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal menghapus data siswa!",
			"status":  500,
		})
		return
	}

	// hapus data user terkait
	if err := config.DB.Model(&user).Where("username", nis).Delete(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal menghapus data user terkait!",
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
