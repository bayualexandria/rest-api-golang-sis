package controllers

import (
	"backend-api/config"
	"backend-api/models"
	siswakelas "backend-api/validations/siswaKelas"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Siswa struct {
	Id           uint64 `json:"id"`
	NIS          string `json:"nis"`
	NamaSiswa    string `json:"nama_siswa"`
	JenisKelamin string `json:"jenis_kelamin"`
	NoHp         string `json:"no_hp"`
	NamaKelas    string `json:"nama_kelas"`
	Jurusan      string `json:"jurusan"`
	TahunAjaran  string `json:"tahun_ajaran"`
	Semester     string `json:"semester"`
	TanggalMasuk string `json:"tanggal_masuk"`
	Status       string `json:"status"`
}

func GetSiswaKelas(c *gin.Context) {
	var data []Siswa

	if err := config.DB.Table("siswa_kelas").
		Joins("JOIN siswa ON siswa_kelas.siswa_id = siswa.id").
		Joins("JOIN kelas ON siswa_kelas.kelas_id = kelas.id").
		Joins("JOIN tahun_ajaran ON siswa_kelas.tahun_ajaran_id = tahun_ajaran.id").
		Joins("JOIN semester ON siswa_kelas.semester_id = semester.id").
		Select("siswa_kelas.id,siswa.nama AS nama_siswa, siswa.nis AS nis, siswa.jenis_kelamin, siswa.no_hp, kelas.nama_kelas, kelas.jurusan, tahun_ajaran.nama_tahun AS tahun_ajaran, semester.nama_semester AS semester, siswa_kelas.tanggal_masuk, siswa_kelas.status").
		Find(&data).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal mengambil data siswa kelas",
			"error":   err.Error(),
			"total":   0,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data siswa kelas berhasil ditampilkan",
		"data":    data,
		"total":   len(data),
	})
}

func GetSiswaKelasByNis(c *gin.Context) {
	idParam := c.Param("nis")
	var data Siswa

	if err := config.DB.Table("siswa_kelas").
		Joins("JOIN siswa ON siswa_kelas.siswa_id = siswa.id").
		Joins("JOIN kelas ON siswa_kelas.kelas_id = kelas.id").
		Joins("JOIN tahun_ajaran ON siswa_kelas.tahun_ajaran_id = tahun_ajaran.id").
		Joins("JOIN semester ON siswa_kelas.semester_id = semester.id").
		Select("siswa_kelas.id,siswa.nis AS nis,siswa.nama AS nama_siswa,  siswa.jenis_kelamin, siswa.no_hp, kelas.nama_kelas, kelas.jurusan, tahun_ajaran.nama_tahun AS tahun_ajaran, semester.nama_semester AS semester, siswa_kelas.tanggal_masuk, siswa_kelas.status").
		Where("siswa.nis = ?", idParam).
		First(&data).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal mengambil data siswa kelas silahkkan periksa NIS",
			"total":   0,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data siswa kelas berhasil ditampilkan",
		"data":    data,
	})
}

func AddSiswaKelas(c *gin.Context) {
	var request = siswakelas.AddDataSiswaKelasRequest{}

	if err := c.ShouldBind(&request); err != nil {
		msg := siswakelas.TranslateAddDataSiswaKelasError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Data tidak valid",
			"error":   msg,
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

	// Cek siswa
	var siswa models.Siswa

	if err := config.DB.Where("id = ?", request.SiswaId).First(&siswa).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Siswa tidak ditemukan",
		})
		return
	}

	// Cek kelas
	var kelas models.Kelas

	if err := config.DB.Where("id = ?", request.KelasId).First(&kelas).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Kelas tidak ditemukan",
		})
		return
	}

	// Cek apakah siswa sudah punya kelas
	var existing models.SiswaKelas

	err := config.DB.
		Table("siswa_kelas sk").
		Joins("JOIN siswa s ON sk.siswa_id = s.id").
		Joins("JOIN tahun_ajaran ta ON sk.tahun_ajaran_id = ta.id").
		Joins("JOIN semester sm ON sk.semester_id = sm.id").
		Where("sk.siswa_id = ?", request.SiswaId).
		Where("sk.tahun_ajaran_id = ?", tahunAjaran.ID).
		Where("sk.semester_id = ?", semester.ID).
		First(&existing).Error

	if err == nil {

		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"message": "Siswa sudah terdaftar di kelas pada tahun ajaran ini",
		})
		return
	}

	if err != nil && err != gorm.ErrRecordNotFound {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal mengecek data siswa",
			"error":   err.Error(),
		})
		return
	}

	// Tanggal masuk
	tanggalMasuk := time.Now()

	// Buat data
	data := models.SiswaKelas{
		SiswaID:       request.SiswaId,
		KelasID:       request.KelasId,
		TahunAjaranID: tahunAjaran.ID,
		SemesterID:    semester.ID,
		TanggalMasuk:  &tanggalMasuk,
		Status:        "aktif",
	}
	// Jika data siswa_id, kelas_id sama dan semester_id dan tahun ajaran_id beda data tersimpan
	if err := config.DB.Where("siswa_id = ? AND kelas_id = ? AND tahun_ajaran_id = ? AND semester_id = ?", request.SiswaId, request.KelasId, tahunAjaran.ID, semester.ID).First(&existing).Error; err == nil {

		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"message": "Siswa sudah terdaftar di kelas ini pada tahun ajaran dan semester yang sama",
		})
		return
	}
	if err := config.DB.Create(&data).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal memasukkan siswa ke kelas",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Siswa berhasil dimasukkan ke kelas",
		"data":    data,
	})
}

func UpdateSiswaKelas(c *gin.Context) {
	id := c.Param("id")
	request := siswakelas.UpdateDataSiswaKelasRequest{}
	var existing models.SiswaKelas

	if err := config.DB.Table("siswa_kelas").Where("id = ?", id).First(&existing).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Data siswa kelas tidak ditemukan",
		})
		return
	}

	if err := c.ShouldBind(&request); err != nil {
		msg := siswakelas.TranslateUpdateDataSiswaKelasError(err)
		c.JSON(400, gin.H{
			"message": "Anda belum merubah data!",
			"data":    msg,
			"status":  400,
		})
		return
	}
	if request.SiswaId != 0 {
		existing.SiswaID = request.SiswaId
	}
	// Apakah bisa jangan pakai 0

	if request.KelasId != 0 {
		existing.KelasID = request.KelasId
	}

	if err := config.DB.Model(&existing).Where("id = ?", id).Updates(&existing).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal mengupdate database: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Data siswa kelas berhasil diupdate",
		"status":  200,
	})

}
