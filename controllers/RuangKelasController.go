package controllers

import (
	"backend-api/config"
	"backend-api/models"
	"net/http"

	ruangKelas "backend-api/validations/ruangKelas"

	"github.com/gin-gonic/gin"
)

type RuangKelasStruct struct {
	Id          uint64 `json:"id"`
	NIP         string `json:"nip" gorm:"column:nip"`
	Name        string `json:"name"`
	KelasId     int    `json:"kelas_id"`
	NamaKelas   string `json:"nama_kelas"`
	Jurusan     string `json:"jurusan"`
	TahunAjaran string `json:"tahun_ajaran"`
	Semester    string `json:"semester"`
	Status      string `json:"status"`
}

func RuangKelas(c *gin.Context) {
	// Implementation for getting room classes

	var data []RuangKelasStruct

	if err := config.DB.Table("wali_kelas").
		Select("wali_kelas.id, guru.nip AS nip, guru.nama AS name, wali_kelas.kelas_id, kelas.nama_kelas, kelas.jurusan, tahun_ajaran.nama_tahun AS tahun_ajaran, semester.nama_semester AS semester, wali_kelas.status").
		Joins("JOIN guru ON wali_kelas.guru_wali_id = guru.nip").
		Joins("JOIN kelas ON wali_kelas.kelas_id = kelas.id").
		Joins("JOIN tahun_ajaran ON wali_kelas.tahun_ajaran_id = tahun_ajaran.id").
		Joins("JOIN semester ON wali_kelas.semester_id = semester.id").
		Find(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal mengambil data ruang kelas",
			"error":   data,
			"total":   0,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data ruang kelas berhasil ditampilkan",
		"data":    data,
		"total":   len(data),
		"status":  200,
	})
}

func RuangKelasById(c *gin.Context) {

	id := c.Param("id")
	var data []RuangKelasStruct

	if err := config.DB.Table("wali_kelas").
		Select("wali_kelas.id, guru.nip AS nip, guru.nama AS name,wali_kelas.kelas_id, kelas.nama_kelas, kelas.jurusan, tahun_ajaran.nama_tahun AS tahun_ajaran, semester.nama_semester AS semester, wali_kelas.status").
		Joins("JOIN guru ON wali_kelas.guru_wali_id = guru.nip").
		Joins("JOIN kelas ON wali_kelas.kelas_id = kelas.id").
		Joins("JOIN tahun_ajaran ON wali_kelas.tahun_ajaran_id = tahun_ajaran.id").
		Joins("JOIN semester ON wali_kelas.semester_id = semester.id").Where("wali_kelas.id = ?", id).
		First(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal mengambil data ruang kelas",
			"error":   data,
			"total":   0,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data ruang kelas berhasil ditampilkan",
		"data":    data,
		"total":   len(data),
		"status":  200,
	})
}

func RuangKelasByNip(c *gin.Context) {
	// Implementation for getting room class by ID
	nip := c.Param("nip")
	var data []RuangKelasStruct

	if err := config.DB.Table("wali_kelas").
		Select("wali_kelas.id, guru.nip AS nip, guru.nama AS name,wali_kelas.kelas_id, kelas.nama_kelas, kelas.jurusan, tahun_ajaran.nama_tahun AS tahun_ajaran, semester.nama_semester AS semester, wali_kelas.status").
		Joins("JOIN guru ON wali_kelas.guru_wali_id = guru.nip").
		Joins("JOIN kelas ON wali_kelas.kelas_id = kelas.id").
		Joins("JOIN tahun_ajaran ON wali_kelas.tahun_ajaran_id = tahun_ajaran.id").
		Joins("JOIN semester ON wali_kelas.semester_id = semester.id").Where("guru.nip = ?", nip).
		Find(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal mengambil data ruang kelas",
			"error":   data,
			"total":   0,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data ruang kelas berhasil ditampilkan",
		"data":    data,
		"total":   len(data),
		"status":  200,
	})
}

func AddRuangKelas(c *gin.Context) {
	var request ruangKelas.AddRuangKelasRequest
	if err := c.ShouldBind(&request); err != nil {
		errors := ruangKelas.TranslateAddRuangKelasError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Validasi gagal",
			"errors":  errors,
			"status":  400,
		})
		return
	}

	// Ambil tahun ajaran aktif
	var tahunAjaran models.TahunAjaran

	if err := config.DB.
		Where("is_active = ?", true).
		First(&tahunAjaran).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Tahun ajaran aktif tidak ditemukan",
		})
		return
	}

	// Ambil semester aktif
	var semester models.Semester

	if err := config.DB.
		Where("is_active = ?", true).
		First(&semester).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Semester aktif tidak ditemukan",
		})
		return
	}

	// Cek apakah wali kelas id sudah ada untuk guru dan kelas yang sama di tahun ajaran yang sama
	var existingWaliKelas models.WaliKelas
	if err := config.DB.
		Where("kelas_id = ?", request.KelasId).
		First(&existingWaliKelas).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Kelas sudah ada dalam tahun ajaran yang sama",
		})
		return
	}

	if err := config.DB.Table("wali_kelas").Create(map[string]interface{}{
		"guru_wali_id":    request.GuruWaliId,
		"kelas_id":        request.KelasId,
		"tahun_ajaran_id": tahunAjaran.Id,
		"semester_id":     semester.Id,
		"status":          "aktif",
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal menambahkan data ruang kelas",
			"error":   err.Error(),
			"status":  500,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data ruang kelas berhasil ditambahkan",
		"status":  200,
	})
}

func UpdateRuangKelas(c *gin.Context) {
	id := c.Param("id")
	var request ruangKelas.UpdateRuangKelasRequest
	var data models.WaliKelas

	if err := config.DB.Table("wali_kelas").Where("id = ?", id).First(&data).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Data ruang kelas tidak ditemukan",
			"status":  404,
		})
		return
	}

	if err := c.ShouldBind(&request); err != nil {
		errors := ruangKelas.TranslateUpdateRuangKelasError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Validasi gagal",
			"errors":  errors,
			"status":  400,
		})
		return
	}

	// 1. Ambil tahun ajaran aktif
	var tahunAjaran models.TahunAjaran
	if err := config.DB.Where("is_active = ?", true).First(&tahunAjaran).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Tahun ajaran aktif tidak ditemukan",
			"status":  400,
		})
		return
	}

	// 2. Ambil semester aktif
	var semester models.Semester
	if err := config.DB.Where("is_active = ?", true).First(&semester).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Semester aktif tidak ditemukan",
			"status":  400,
		})
		return
	}

	// 3. Cek apakah ada siswa di kelas tersebut
	var totalSiswa int64
	err := config.DB.Model(&models.SiswaKelas{}).
		Where("kelas_id = ? AND tahun_ajaran_id = ? AND semester_id = ?", data.KelasId, tahunAjaran.Id, semester.Id).
		Count(&totalSiswa).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal memeriksa data siswa kelas",
			"status":  500,
		})
		return
	}

	// Jika jumlah siswa > 0, tolak proses update
	if totalSiswa > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Tidak bisa mengubah data kelas dikarenakan ada data siswa yang sudah masuk dalam kelas tersebut!",
			"status":  400,
		})
		return
	}

	// 4. Lakukan update data wali_kelas
	if err := config.DB.Table("wali_kelas").Where("id = ?", id).Updates(map[string]interface{}{
		"guru_wali_id":    request.GuruWaliId,
		"kelas_id":        request.KelasId,
		"tahun_ajaran_id": tahunAjaran.Id,
		"semester_id":     semester.Id,
		"status":          "aktif",
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal mengupdate data ruang kelas",
			"error":   err.Error(),
			"status":  500,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data ruang kelas berhasil diupdate",
		"status":  200,
	})
}
