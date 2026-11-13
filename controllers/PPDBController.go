package controllers

import (
	"backend-api/config"
	"backend-api/models"
	"backend-api/validations/ppdb"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserCalonSiswaByUsername(c *gin.Context) {
	username := c.Param("username")
	var userCalonSiswa models.UserCalonSiswa

	if err := config.DB.Where("username = ?", username).First(&userCalonSiswa).Error; err != nil {
		c.JSON(404, gin.H{"message": "User calon siswa tidak ditemukan", "status": 404})
		return
	}

	c.JSON(200, gin.H{"message": "Berhasil mendapatkan data user calon siswa", "status": 200, "data": userCalonSiswa})
}

type DataCalonSiswaResponse struct {
	NamaLengkap  string `json:"nama_lengkap"`
	NoRegistrasi string `json:"no_registrasi"`
	Email        string `json:"email"`
	Nik          string `json:"nik"`
	Nisn         string `json:"nisn"`
	JenisKelamin string `json:"jenis_kelamin"`
	Agama        string `json:"agama"`
	Ttl          string `json:"ttl"`
	Alamat       string `json:"alamat"`
	SekolahAsal  string `json:"sekolah_asal"`
	TahunLulus   string `json:"tahun_lulus"`
	// Data Orang Tua
	NikAyah       string `json:"nik_ayah"`
	NamaAyah      string `json:"nama_ayah"`
	PekerjaanAyah string `json:"pekerjaan_ayah"`
	NikIbu        string `json:"nik_ibu"`
	NamaIbu       string `json:"nama_ibu"`
	PekerjaanIbu  string `json:"pekerjaan_ibu"`
	NamaWali      string `json:"nama_wali"`
	PekerjaanWali string `json:"pekerjaan_wali"`
	NomorTelepon  string `json:"nomor_telepon"`
	Status        string `json:"status"`
}

func GetDataCalonSiswa(c *gin.Context) {
	var getData []DataCalonSiswaResponse

	if err := config.DB.Table("user_calon_siswa").
		Select("user_calon_siswa.nama_lengkap, user_calon_siswa.username AS no_registrasi, user_calon_siswa.email, calon_siswa.nik, calon_siswa.nisn, calon_siswa.jenis_kelamin, calon_siswa.agama, calon_siswa.ttl, calon_siswa.alamat, calon_siswa.sekolah_asal, calon_siswa.tahun_lulus, orang_tua.nik_ayah, orang_tua.nama_ayah, orang_tua.pekerjaan_ayah, orang_tua.nik_ibu, orang_tua.nama_ibu, orang_tua.pekerjaan_ibu, orang_tua.nama_wali, orang_tua.pekerjaan_wali, orang_tua.nomor_telepon, pendaftaran.status_verifikasi AS status").
		Joins("JOIN orang_tua ON user_calon_siswa.username = orang_tua.id_siswa").
		Joins("JOIN calon_siswa ON user_calon_siswa.username = calon_siswa.id_user_calon_siswa").
		Joins("JOIN pendaftaran ON user_calon_siswa.username = pendaftaran.nomor_pendaftaran").
		Scan(&getData).Error; err != nil {
		c.JSON(500, gin.H{"message": "Gagal mendapatkan data calon siswa", "status": 500})
		return
	}
	c.JSON(200, gin.H{"message": "Berhasil mendapatkan data calon siswa", "status": 200, "data": getData})
}

func AddDataCalonSiswa(c *gin.Context) {
	var request ppdb.AddDataFormulirCalonSiswaRequest

	if err := c.ShouldBind(&request); err != nil {
		msg := ppdb.GetAddDataFormulirCalonSiswaValidationError(err)
		c.JSON(400, gin.H{"message": msg, "status": 400})
		return
	}

	// Ambil file multipart dengan key "dokumen"
	form, err := c.MultipartForm()
	var files []*multipart.FileHeader
	if err == nil && form != nil {
		files = form.File["dokumen"]
	}

	// Gunakan Database Transaction
	errTx := config.DB.Transaction(func(tx *gorm.DB) error {
		// Simpan Data Calon Siswa
		calonSiswa := models.CalonSiswa{
			NamaLengkap:      request.NamaLengkap,
			IdUserCalonSiswa: request.IdUserCalonSiswa,
			Nik:              request.Nik,
			Nisn:             request.Nisn,
			JenisKelamin:     request.JenisKelamin,
			Agama:            request.Agama,
			Ttl:              request.Ttl,
			Alamat:           request.Alamat,
			SekolahAsal:      request.SekolahAsal,
			TahunLulus:       request.TahunLulus,
		}

		if err := tx.Create(&calonSiswa).Error; err != nil {
			return err
		}

		// Simpan Data Orang Tua
		orangTua := models.OrangTua{
			IdSiswa:       request.IdUserCalonSiswa,
			NikAyah:       request.NikAyah,
			NamaAyah:      request.NamaAyah,
			PekerjaanAyah: request.PekerjaanAyah,
			NikIbu:        request.NikIbu,
			NamaIbu:       request.NamaIbu,
			PekerjaanIbu:  request.PekerjaanIbu,
			NamaWali:      request.NamaWali,
			PekerjaanWali: request.PekerjaanWali,
			NomorTelepon:  request.NomorTelepon,
		}

		if err := tx.Create(&orangTua).Error; err != nil {
			return err
		}

		// ----------------------------------------------------
		// UBAH DIREKTORI PENYIMPANAN MENJADI: storage/ppdb/{id_user_calon_siswa}
		// ----------------------------------------------------
		uploadDir := filepath.Join("storage", "ppdb", request.IdUserCalonSiswa)
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			return err
		}

		// Looping file dokumen yang diunggah
		for i, file := range files {
			// Tentukan jenis dokumen sesuai inputan (contoh: 'Ijazah', 'Kartu Keluarga', dll)
			jenisDokumen := "Lainnya"
			if len(request.JenisDokumen) > i && request.JenisDokumen[i] != "" {
				jenisDokumen = request.JenisDokumen[i]
			}

			// Buat nama file unik
			ext := filepath.Ext(file.Filename)
			uniqueName := fmt.Sprintf("%s_%s_%d%s", request.Nisn, jenisDokumen, time.Now().UnixNano(), ext)
			dst := filepath.Join(uploadDir, uniqueName)

			// Simpan file fisik ke server
			if err := c.SaveUploadedFile(file, dst); err != nil {
				return err
			}

			// Simpan ke tabel dokumen_persyaratan_calon_siswa
			dokumen := models.DokumenPersyaratanCalonSiswa{
				IdUserCalonSiswa: request.IdUserCalonSiswa,
				JenisDokumen:     jenisDokumen,
				NamaFile:         file.Filename,
				PathFile:         dst, // Contoh hasil path: storage/ppdb/user123/0081234567_Ijazah_171...pdf
				DiunggahPada:     time.Now(),
			}

			if err := tx.Create(&dokumen).Error; err != nil {
				return err
			}
		}

		return nil
	})

	var calonSiswa models.CalonSiswa

	if err := config.DB.Model(&calonSiswa).Where("id_user_calon_siswa = ?", request.IdUserCalonSiswa).First(&calonSiswa).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal menyimpan data pendaftaran dan dokumen persyaratan",
			"error":   errTx.Error(),
			"status":  500,
		})
		return

	}

	if errTx != nil {
		c.JSON(500, gin.H{
			"message": "Gagal menyimpan data pendaftaran dan dokumen persyaratan",
			"error":   errTx.Error(),
			"status":  500,
		})
		return
	}

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

	if err := config.DB.Table("pendaftaran").Create(map[string]interface{}{
		"nomor_pendaftaran":    request.IdUserCalonSiswa,
		"id_jalur_pendaftaran": 1,
		"tanggal_daftar":       tanggalDB,
		"status_verifikasi":    "Pending",
		"catatan_verifikasi":   "",
		"status_kelulusan":     "Proses",
	}).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal menyimpan data pendaftaran dan dokumen persyaratan",
			"status":  500,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Data calon siswa beserta dokumen persyaratan berhasil disimpan",
		"status":  200,
	})
}

func GetDokumenPersyaratanCalonSiswa(c *gin.Context) {
	username := c.Param("username")
	var persyaratan []models.DokumenPersyaratanCalonSiswa

	if err := config.DB.Where("id_user_calon_siswa = ?", username).Find(&persyaratan).Error; err != nil {
		c.JSON(404, gin.H{"message": "Dokumen persyaratan calon siswa tidak ditemukan", "status": 404})
		return
	}
	c.JSON(200, gin.H{"message": "Berhasil mendapatkan data dokumen persyaratan calon siswa", "status": 200, "data": persyaratan})
}

func GetDataCalonSiswaByUsername(c *gin.Context) {
	username := c.Param("username")
	var calonSiswa models.CalonSiswa
	var orangTua models.OrangTua

	if err := config.DB.Where("id_user_calon_siswa = ?", username).First(&calonSiswa).Error; err != nil {
		c.JSON(404, gin.H{"message": "Data calon siswa tidak ditemukan", "status": 404})
		return
	}

	if err := config.DB.Where("id_siswa = ?", username).First(&orangTua).Error; err != nil {
		c.JSON(404, gin.H{"message": "Data orang tua calon siswa tidak ditemukan", "status": 404})
		return
	}

	c.JSON(200, gin.H{
		"message": "Berhasil mendapatkan data calon siswa dan orang tua",
		"status":  200,
		"data": gin.H{
			"calon_siswa": calonSiswa,
			"orang_tua":   orangTua,
		},
	})
}

func GetDataPPDBSiswaByUsername(c *gin.Context) {
	username := c.Param("username")
	var calonSiswa models.CalonSiswa
	var orangTua models.OrangTua

	if err := config.DB.Where("id_user_calon_siswa = ?", username).First(&calonSiswa).Error; err != nil {
		c.JSON(404, gin.H{"message": "Data calon siswa tidak ditemukan", "status": 404})
		return
	}

	if err := config.DB.Where("id_siswa = ?", username).First(&orangTua).Error; err != nil {
		c.JSON(404, gin.H{"message": "Data orang tua calon siswa tidak ditemukan", "status": 404})
		return
	}

	c.JSON(200, gin.H{
		"message": "Berhasil mendapatkan data calon siswa dan orang tua",
		"status":  200,
		"data": gin.H{
			"calon_siswa": calonSiswa,
			"orang_tua":   orangTua,
		},
	})
}

func GetDokumenPersyaratanDataCalonSiswa(c *gin.Context) {
	username := c.Param("username")
	var persyaratan []models.DokumenPersyaratanCalonSiswa

	if err := config.DB.Where("id_user_calon_siswa = ?", username).Find(&persyaratan).Error; err != nil {
		c.JSON(404, gin.H{"message": "Dokumen persyaratan calon siswa tidak ditemukan", "status": 404})
		return
	}
	c.JSON(200, gin.H{"message": "Berhasil mendapatkan data dokumen persyaratan calon siswa", "status": 200, "data": persyaratan})
}
