package controllers

import (
	"backend-api/config"

	"github.com/gin-gonic/gin"
)

type DataPenilaian struct {
	Id                 uint64 `json:"id"`
	NamaKelas          string `json:"nama_kelas"`
	Jurusan            string `json:"jurusan"`
	KodeMapel          string `json:"kode_mapel"`
	NamaMapel          string `json:"nama_mapel"`
	NamaGuru           string `json:"nama_guru"`
	Nip                string `json:"nip"`
	NamaJenisPenilaian string `json:"nama_jenis_penilaian"`
	NamaSemester       string `json:"nama_semester"`
}

func GetDataPenilaian(c *gin.Context) {
	var penilaian []DataPenilaian

	if err := config.DB.Table("penilaian").
		Joins("JOIN kelas ON penilaian.kelas_id = kelas.id").
		Joins("JOIN mata_pelajaran ON penilaian.mata_pelajaran_id = mata_pelajaran.id").
		Joins("JOIN guru ON penilaian.guru_id = guru.id").
		Joins("JOIN jenis_penilaian ON penilaian.jenis_penilaian_id=jenis_penilaian.id").
		Joins("JOIN semester ON penilaian.semester_id=semester.id").Select("").Scan(&penilaian).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal mengambil data guru",
			"status":  500,
		})
		return
	}

}
