package controllers

import (
	"backend-api/config"
	"backend-api/models"
	absensisiswa "backend-api/validations/absensiSiswa"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func AddAbsensiSiswa(c *gin.Context) {

	var request absensisiswa.AddAbsensiValidation

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Data absensi tidak valid",
			"error":   absensisiswa.TranslateAddAbsensiError(err),
		})

		return
	}

	// =====================================================
	// Validasi status kehadiran
	// =====================================================

	var statusKehadiran models.StatusKehadiran

	if err := config.DB.
		First(&statusKehadiran, request.StatusKehadiranID).Error; err != nil {

		if err == gorm.ErrRecordNotFound {

			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Status kehadiran tidak ditemukan",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal mengambil status kehadiran",
			"error":   err.Error(),
		})

		return
	}

	// =====================================================
	// Cari siswa kelas
	// =====================================================

	var siswaKelas models.SiswaKelas

	if err := config.DB.
		Preload("Siswa").
		Preload("Kelas").
		Preload("TahunAjaran").
		First(&siswaKelas, request.SiswaKelasID).Error; err != nil {

		if err == gorm.ErrRecordNotFound {

			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Data siswa kelas tidak ditemukan",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal mengambil data siswa kelas",
			"error":   err.Error(),
		})

		return
	}

	// =====================================================
	// Pastikan siswa masih aktif
	// =====================================================

	if siswaKelas.Status != "aktif" {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Siswa tidak berstatus aktif pada kelas ini",
		})

		return
	}

	// =====================================================
	// Cari semester aktif
	// =====================================================

	var semester models.Semester

	if err := config.DB.
		Where("tahun_ajaran_id = ?", siswaKelas.TahunAjaranID).
		Where("is_active = ?", true).
		First(&semester).Error; err != nil {

		if err == gorm.ErrRecordNotFound {

			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Semester aktif tidak ditemukan",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal mengambil semester aktif",
			"error":   err.Error(),
		})

		return
	}

	// =====================================================
	// Tanggal hari ini
	// =====================================================

	now := time.Now()

	tanggalDB := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0,
		0,
		0,
		0,
		now.Location(),
	)

	// =====================================================
	// Cek apakah sudah absen hari ini
	// =====================================================

	var existing models.AbsensiSiswa

	err := config.DB.
		Where("siswa_kelas_id = ?", request.SiswaKelasID).
		Where("tanggal = ?", tanggalDB).
		First(&existing).Error

	if err == nil {

		// Ambil relasi status kehadiran
		config.DB.
			Preload("StatusKehadiran").
			First(&existing, existing.ID)

		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"message": "Siswa sudah memiliki absensi hari ini",
			"data":    existing,
		})

		return
	}

	if err != nil && err != gorm.ErrRecordNotFound {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal mengecek absensi",
			"error":   err.Error(),
		})

		return
	}

	// =====================================================
	// Buat data absensi
	// =====================================================

	data := models.AbsensiSiswa{
		SiswaKelasID:      request.SiswaKelasID,
		SemesterID:        semester.ID,
		StatusKehadiranID: request.StatusKehadiranID,
		Tanggal:           tanggalDB,
		Keterangan:        request.Keterangan,

		// Default NULL
		JamMasuk:  nil,
		JamKeluar: nil,
	}

	// =====================================================
	// Jika status = HADIR
	// Catat jam masuk
	// =====================================================

	if statusKehadiran.Kode == "H" {

		jamMasuk := time.Now()

		data.JamMasuk = &jamMasuk
	}

	// =====================================================
	// Simpan absensi
	// =====================================================

	if err := config.DB.Create(&data).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal menyimpan absensi",
			"error":   err.Error(),
		})

		return
	}



	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Absensi berhasil disimpan",
		"data":    data,
	})
}
