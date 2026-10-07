package controllers

import (
	"backend-api/config"
	"backend-api/models"
	gurumapel "backend-api/validations/guruMapel"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GuruMapelResponse struct {
	Id          int64  `json:"id" gorm:"column:id"`
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
		guru_mata_pelajaran.id,
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
		"status":  200,
	})
}

func GetDataGuruMapelById(c *gin.Context) {
	id := c.Param("id")
	var guruMapel models.GuruMapel

	if err := config.DB.Model(&guruMapel).Where("id = ?", id).First(&guruMapel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Guru mapel tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    guruMapel,
		"message": "Data berhasil ditampilkan!",
		"status":  200,
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

func UpdateGuruMapel(c *gin.Context) {
	var request gurumapel.UpdateGuruMapelRequest
	var guruMapel models.GuruMapel

	id := c.Param("id")

	if err := config.DB.Model(&guruMapel).Where("id = ?", id).First(&guruMapel).Error; err != nil {
		c.JSON(404, gin.H{
			"message": "Data guru mapel dengan id " + id + " tidak ditemukan",
			"status":  404,
		})

		return
	}

	if err := c.ShouldBind(&request); err != nil {
		msg := gurumapel.TranslateUpdatedGuruMapelError(err)
		c.JSON(400, gin.H{
			"message": "Anda belum merubah data!",
			"data":    msg,
			"status":  400,
		})

		return
	}

	var guru models.Guru
	var mapel models.Mapel
	var tahunAjaran models.TahunAjaran

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

	if err := config.DB.Model(&guruMapel).Where("id = ?", id).Updates(map[string]interface{}{
		"guru_id":           request.GuruId,
		"mata_pelajaran_id": request.MataPelajaranId,
		"tahun_ajaran_id":   tahunAjaran.Id,
	}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Gagal mengupdate data guru mapel!",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"guru_mapel":     guru.Nama,
			"mata_pelajaran": mapel.Nama,
			"tahun_ajaran":   tahunAjaran.NamaTahun,
		},
		"message": "Berhasil mengubah data guru mapel",
		"status":  200,
	})
}

func DeleteGuruMapel(c *gin.Context) {
	var guruMapel models.GuruMapel
	id := c.Param("id")

	if err := config.DB.Model(&guruMapel).Where("id = ?", id).First(&guruMapel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Guru mata pelajaran tidak ditemukan",
		})
		return
	}

	if err := config.DB.Model(&guruMapel).Where("id = ?", id).Delete(&guruMapel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Gagal menghapus data guru mapel!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Berhasil menghapus data guru mapel",
	})
}
