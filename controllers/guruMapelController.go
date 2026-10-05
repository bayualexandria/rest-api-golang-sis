package controllers

import (
	"backend-api/config"
	"backend-api/models"
	gurumapel "backend-api/validations/guruMapel"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GuruMapelResponse struct {
	NIP         string `json:"nip" gorm:"column:nip"`
	Nama        string `json:"nama" gorm:"column:nama"`
	KodeMapel   string `json:"kode_mapel" gorm:"column:kode_mapel"`
	NamaMapel   string `json:"nama_mapel" gorm:"column:nama_mapel"`
	TahunAjaran string `json:"tahun_ajaran" gorm:"column:tahun_ajaran"`
}

func GetAllGuruMapel(c *gin.Context) {
	// Inisialisasi slice agar output JSON berupa [] bukan null saat kosong
	guruMapel := make([]GuruMapelResponse, 0)

	err := config.DB.
		Table("guru_mata_pelajaran").
		Select(`
			guru.nip AS nip,
			guru.nama AS nama,
			mata_pelajaran.kode AS kode_mapel,
			mata_pelajaran.nama AS nama_mapel,
			tahun_ajaran.nama_tahun AS tahun_ajaran
		`).
		Joins("LEFT JOIN guru ON guru_mata_pelajaran.guru_id = guru.id").
		Joins("LEFT JOIN mata_pelajaran ON guru_mata_pelajaran.mata_pelajaran_id = mata_pelajaran.id").
		Joins("LEFT JOIN tahun_ajaran ON guru_mata_pelajaran.tahun_ajaran_id = tahun_ajaran.id").
		Scan(&guruMapel).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Gagal mendapatkan data guru mapel",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    guruMapel,
		"message": "Berhasil mendapatkan data guru mapel",
	})
}

func AddGuruMapel(c *gin.Context) {
	var request gurumapel.AddGuruMapelRequest
	var guruMapel models.GuruMapel
	var guru models.Guru
	var mapel models.Mapel
	var tahunAjaran models.TahunAjaran

	if err := c.ShouldBind(&request); err != nil {
		msg := gurumapel.TranslateAddGuruMapelError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"errors":  msg,
			"message": "Validasi gagal",
		})
		return
	}

	// Cek apakah guru dengan NIP yang diberikan ada
	if err := config.DB.First(&guru, "id = ?", request.GuruId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Guru tidak ditemukan",
		})
		return
	}

	// Cek apakah mata pelajaran dengan kode yang diberikan ada
	if err := config.DB.First(&mapel, "id = ?", request.MataPelajaranId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Mata pelajaran tidak ditemukan",
		})
		return
	}

	// Cek apakah tahun ajaran dengan ID yang diberikan ada
	if err := config.DB.
		Where("is_active = ?", true).
		First(&tahunAjaran).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Tahun ajaran aktif tidak ditemukan",
		})
		return
	}

	// Cek apakah mapel id sudah ada untuk guru yang sama di tahun ajaran yang sama
	var existingGuruMapel models.GuruMapel
	if err := config.DB.
		Where("guru_id = ? AND mata_pelajaran_id = ?", request.GuruId, request.MataPelajaranId).
		First(&existingGuruMapel).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Guru sudah mengajar mata pelajaran ini dalam tahun ajaran yang sama",
		})
		return
	}

	if err := config.DB.Model(&guruMapel).Create(map[string]interface{}{
		"guru_id":           request.GuruId,
		"mata_pelajaran_id": request.MataPelajaranId,
		"tahun_ajaran_id":   tahunAjaran.Id,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Gagal menambahkan data guru mapel",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    guruMapel,
		"message": "Berhasil menambahkan data guru mapel",
	})
}
