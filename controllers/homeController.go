package controllers

import (
	"backend-api/config"
	"backend-api/models"
	sekolahController "backend-api/validations/sekolah"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"data":    []string{},
		"message": "Data berhasil ditampilkan!",
		"success": true,
	})
}

func ProfileSekolahHandler(c *gin.Context) {
	var sekolah models.ProfileSekolah
	err := config.DB.First(&sekolah).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal mengambil data profile sekolah",
			"status":  500,
		})
		return
	}
	c.JSON(200, gin.H{
		"data":    sekolah,
		"message": "Data berhasil ditampilkan!",
		"success": true,
	})
}

func UpdateProfileSekolahHandler(c *gin.Context) {
	var sekolah models.ProfileSekolah
	var input sekolahController.UpdateDataSekolahValidation

	if err := config.DB.First(&sekolah).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal mengambil data profile sekolah",
			"status":  500,
		})
		return
	}

	if err := c.ShouldBind(&input); err != nil {
		msg := sekolahController.TranslateUpdateDataSekolahError(err)
		c.JSON(400, gin.H{
			"message": "Data input tidak valid",
			"error":   msg,
			"status":  400,
		})
		return
	}

	// =========================================================
	// UPDATE FOTO PROFILE
	// =========================================================

	if input.ImageProfile != nil {
		file := input.ImageProfile

		// Folder penyimpanan berdasarkan NIP
		folderPath := filepath.Join(
			"storage",
			"sekolah",
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

		oldImagePath := filepath.Clean(sekolah.ImageProfile)

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
		sekolah.ImageProfile = filePath
	}

	if input.NamaSekolah != "" {
		sekolah.NamaSekolah = input.NamaSekolah
	}

	if input.Alamat != "" {
		sekolah.Alamat = input.Alamat
	}

	if input.NoTelp != "" {
		sekolah.NoTelp = input.NoTelp
	}

	if input.Akreditasi != "" {
		sekolah.Akreditasi = input.Akreditasi
	}

	if err := config.DB.Save(&sekolah).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal memperbarui data profile sekolah",
			"status":  500,
		})
		return
	}

	c.JSON(200, gin.H{
		"data":    sekolah,
		"message": "Data berhasil diperbarui!",
		"success": true,
	})

}
