package controllers

import (
	"backend-api/config"
	"backend-api/models"
	jadwalpelajaran "backend-api/validations/jadwalPelajaran"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetDataJadwalPelajaranByKelasStruct struct {
	Id         uint64 `json:"id" gorm:"column:id"`
	KelasId    int64  `json:"kelas_id" gorm:"column:kelas_id"`
	NamaKelas  string `json:"nama_kelas" gorm:"column:nama_kelas"`
	Jurusan    string `json:"jurusan" gorm:"column:jurusan"`
	NamaGuru   string `json:"nama_guru" gorm:"column:nama_guru"`
	NamaMapel  string `json:"nama_mapel" gorm:"column:nama_mapel"`
	NamaTahun  string `json:"nama_tahun" gorm:"column:nama_tahun"`
	Hari       string `json:"hari" gorm:"column:hari"`
	JamMulai   string `json:"jam_mulai" gorm:"column:jam_mulai"`     // Ubah ke string jika di DB tipe datanya TIME/VARCHAR
	JamSelesai string `json:"jam_selesai" gorm:"column:jam_selesai"` // Ubah ke string
}

func GetJadwalPelajaranByKelas(c *gin.Context) {
	var dataJadwalMapel []GetDataJadwalPelajaranByKelasStruct
	kelasId := c.Param("id")

	err := config.DB.Table("jadwal_pelajaran").
		Joins("JOIN guru ON jadwal_pelajaran.guru_id = guru.id").
		Joins("JOIN kelas ON jadwal_pelajaran.kelas_id = kelas.id").
		Joins("JOIN mata_pelajaran ON jadwal_pelajaran.mata_pelajaran_id = mata_pelajaran.id").
		Joins("JOIN tahun_ajaran ON jadwal_pelajaran.tahun_ajaran_id = tahun_ajaran.id").
		Select("jadwal_pelajaran.id, kelas.id AS kelas_id, kelas.nama_kelas, kelas.jurusan, guru.nama AS nama_guru, mata_pelajaran.nama AS nama_mapel, tahun_ajaran.nama_tahun, jadwal_pelajaran.hari, jadwal_pelajaran.jam_mulai, jadwal_pelajaran.jam_selesai").
		Where("jadwal_pelajaran.kelas_id = ?", kelasId).
		Scan(&dataJadwalMapel).Error

	// Perbaikan 1: Cek jika terjadi error database
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Terjadi kesalahan pada server!",
			"error":   err.Error(),
			"status":  500,
		})
		return
	}

	// Perbaikan 2: Cek jika data kosong
	if len(dataJadwalMapel) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Jadwal mapel tidak ditemukan!",
			"status":  404,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dataJadwalMapel, "status": 200})
}

func GetJadwalPelajaranByGuruMapel(c *gin.Context) {
	var dataJadwalMapel []GetDataJadwalPelajaranByKelasStruct
	guruId := c.Param("idGuru")
	mapelId := c.Param("idMapel")

	err := config.DB.Table("jadwal_pelajaran").
		Joins("JOIN guru ON jadwal_pelajaran.guru_id = guru.id").
		Joins("JOIN kelas ON jadwal_pelajaran.kelas_id = kelas.id").
		Joins("JOIN mata_pelajaran ON jadwal_pelajaran.mata_pelajaran_id = mata_pelajaran.id").
		Joins("JOIN tahun_ajaran ON jadwal_pelajaran.tahun_ajaran_id = tahun_ajaran.id").
		Select("jadwal_pelajaran.id, kelas.id AS kelas_id, kelas.nama_kelas, kelas.jurusan, guru.nama AS nama_guru, mata_pelajaran.nama AS nama_mapel, tahun_ajaran.nama_tahun, jadwal_pelajaran.hari, jadwal_pelajaran.jam_mulai, jadwal_pelajaran.jam_selesai").
		Where("jadwal_pelajaran.guru_id = ? AND jadwal_pelajaran.mata_pelajaran_id = ?", guruId, mapelId).
		Scan(&dataJadwalMapel).Error

	// 1. Cek jika terjadi error pada database (gunakan err != nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Terjadi kesalahan pada server!",
			"error":   err.Error(),
			"status":  500,
		})
		return
	}

	// 2. Cek jika data tidak ditemukan (slice kosong)
	if len(dataJadwalMapel) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Jadwal mapel tidak ditemukan!",
			"status":  404,
		})
		return
	}

	// 3. Jika berhasil dan data ada
	c.JSON(http.StatusOK, gin.H{
		"data":   dataJadwalMapel,
		"status": 200,
	})
}

func InsertJadwalPelajaran(c *gin.Context) {
	var input jadwalpelajaran.AddJadwalPelajaran

	if err := c.ShouldBind(&input); err != nil {
		errors := jadwalpelajaran.TranslateAddJadwalPelajaranError(err)
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

	// Ambil data guru
	var guru models.Guru

	if err := config.DB.Where("id = ?", input.GuruId).First(&guru).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Data guru tidak ditemukan!",
		})
		return
	}

	// Ambil data kelas
	var kelas models.Kelas

	if err := config.DB.Where("id = ?", input.KelasId).First(&kelas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Data kelas tidak ditemukan!",
		})
		return
	}

	// Ambil data mata pelajaran
	var mapel models.Mapel

	if err := config.DB.Where("id = ?", input.MataPelajaranId).First(&mapel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Data mata pelajaran tidak ditemukan!",
		})
		return
	}

	// Cek apakah jadwal pelajaran hari sudah ada untuk guru dan kelas yang sama di tahun ajaran yang sama
	var existingWaliKelas models.JadwalPelajaran
	var count int64
	err := config.DB.Model(existingWaliKelas).
		Where("kelas_id = ?", input.KelasId). // Filter berdasarkan kelas
		Where("hari = ?", input.Hari).
		Where("jam_mulai < ? AND jam_selesai > ?", input.JamSelesai, input.JamMulai).
		Count(&count).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal memeriksa jadwal",
		})
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Jadwal pada kelas ini bentrok dengan jam pelajaran lain",
		})
		return
	}

	if err := config.DB.Table("jadwal_pelajaran").Create(map[string]interface{}{
		"kelas_id":          input.KelasId,
		"guru_id":           input.GuruId,
		"mata_pelajaran_id": input.MataPelajaranId,
		"tahun_ajaran_id":   tahunAjaran.Id,
		"hari":              input.Hari,
		"jam_mulai":         input.JamMulai,
		"jam_selesai":       input.JamSelesai,
	}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Gagal menyimpan data jadwal pelajaran!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data jadwal pelajaran berhasil ditambahkan",
		"status":  200,
	})

}

func UpdateJadwalPelajaran(c *gin.Context) {
	id := c.Param("id")
	var input jadwalpelajaran.UpdateJadwalPelajaran

	if err := c.ShouldBind(&input); err != nil {
		errors := jadwalpelajaran.TranslateAddJadwalPelajaranError(err)
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

	// Ambil data guru
	var guru models.Guru

	if err := config.DB.Where("id = ?", input.GuruId).First(&guru).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Data guru tidak ditemukan!",
		})
		return
	}

	// Ambil data kelas
	var kelas models.Kelas

	if err := config.DB.Where("id = ?", input.KelasId).First(&kelas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Data kelas tidak ditemukan!",
		})
		return
	}

	// Ambil data mata pelajaran
	var mapel models.Mapel

	if err := config.DB.Where("id = ?", input.MataPelajaranId).First(&mapel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Data mata pelajaran tidak ditemukan!",
		})
		return
	}

	// Cek apakah jadwal pelajaran hari sudah ada untuk guru dan kelas yang sama di tahun ajaran yang sama
	var existingWaliKelas models.JadwalPelajaran
	var count int64
	err := config.DB.Model(existingWaliKelas).
		Where("kelas_id = ?", input.KelasId). // Filter berdasarkan kelas
		Where("hari = ?", input.Hari).
		Where("jam_mulai < ? AND jam_selesai > ?", input.JamSelesai, input.JamMulai).
		Count(&count).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal memeriksa jadwal",
		})
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Jadwal pada kelas ini bentrok dengan jam pelajaran lain",
		})
		return
	}

	if err := config.DB.Table("jadwal_pelajaran").Where("id = ?", id).Updates(map[string]interface{}{
		"kelas_id":          input.KelasId,
		"guru_id":           input.GuruId,
		"mata_pelajaran_id": input.MataPelajaranId,
		"tahun_ajaran_id":   tahunAjaran.Id,
		"hari":              input.Hari,
		"jam_mulai":         input.JamMulai,
		"jam_selesai":       input.JamSelesai,
	}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Gagal mengupdate data jadwal pelajaran!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data jadwal pelajaran berhasil diubah!",
		"status":  200,
	})

}
