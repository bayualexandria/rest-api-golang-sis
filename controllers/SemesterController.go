package controllers

import (
	"backend-api/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SemesterWithTahunAjaran struct {
	Id           uint   `json:"id"`
	NamaSemester string `json:"nama_semester"`
	NamaTahun    string `json:"nama_tahun"`
}

func GetSemester(c *gin.Context) {
	var semesters SemesterWithTahunAjaran
	// Menampilkan nama_semeseter dan nama_tahun ajaran dari tabel semester dan tahun_ajaran
	config.DB.Table("semester").
		Select("semester.id, semester.nama_semester,  tahun_ajaran.nama_tahun").
		Joins("JOIN tahun_ajaran ON semester.tahun_ajaran_id = tahun_ajaran.id").Find(&semesters)
	c.JSON(http.StatusOK, gin.H{"data": semesters})
}
