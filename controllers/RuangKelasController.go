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
	if err := config.DB.Table("wali_kelas").Create(map[string]interface{}{
		"guru_wali_id":    request.GuruWaliId,
		"kelas_id":        request.KelasId,
		"tahun_ajaran_id": tahunAjaran.ID,
		"semester_id":     semester.ID,
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
