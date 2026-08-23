package controllers

import (
	"backend-api/config"
	"backend-api/models"
	"time"

	absensisiswa "backend-api/validations/absensiSiswa"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AbsensiTampilData struct {
	Name         string `json:name`
	Nis          int    `json:nis`
	NamaKelas    string `json:nama_kelas`
	Jurusan      string `json:jurusan`
	ImageProfile string `json:image_profile`
}

func AddAbsensiSiswa(c *gin.Context) {

	var request absensisiswa.AddAbsensiValidation

	if err := c.ShouldBind(&request); err != nil {

		msg := absensisiswa.TranslateAddAbsensiError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": msg,
		})

		return
	}

	// =====================================================
	// Validasi status kehadiran
	// =====================================================

	var siswa models.Siswa

	if err := config.DB.Where("nis = ?", request.Nis).First(&siswa).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Data siswa tidak ditemukan",
		})

		return
	}

	// =====================================================
	// Cari siswa kelas
	// =====================================================

	var siswaKelas models.SiswaKelas

	if err := config.DB.Where("siswa_id =?", siswa.Id).First(&siswaKelas).Error; err != nil {

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

	if err := config.DB.
		Where("siswa_kelas_id = ?", siswaKelas.ID).
		Where("tanggal = ?", tanggalDB).
		First(&existing).Error; err == nil {

		// Ambil relasi status kehadiran
		config.DB.Where("id = ?", existing.ID).First(&existing)

		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"message": "Siswa sudah memiliki absensi hari ini",
			"data":    existing,
		})

		return
	}

	// =====================================================
	// Buat data absensi
	// =====================================================
	jamMasuk := time.Now()
	data := models.AbsensiSiswa{
		SiswaKelasID:      siswaKelas.ID,
		SemesterID:        semester.ID,
		StatusKehadiranID: 1,
		Tanggal:           tanggalDB,
		Keterangan:        request.Keterangan,

		// Default NULL
		JamMasuk:  &jamMasuk,
		JamKeluar: nil,
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

	var kelas models.Kelas
	if err := config.DB.Model(&kelas).Where("id =?", siswaKelas.KelasID).First(&kelas).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Data kelas tidak ditemukan",
		})

		return
	}
	dataTampil := AbsensiTampilData{
		Name:         siswa.Nama,
		Nis:          siswa.Nis,
		NamaKelas:    kelas.NamaKelas,
		Jurusan:      kelas.Jurusan,
		ImageProfile: siswa.ImageProfile,
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Absensi berhasil disimpan",
		"data":    dataTampil,
	})
}
