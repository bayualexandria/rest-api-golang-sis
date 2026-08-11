package services

import (
	"backend-api/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type SemesterService struct {
	DB *gorm.DB
}

func NewSemesterService(db *gorm.DB) *SemesterService {
	return &SemesterService{
		DB: db,
	}
}

func (s *SemesterService) EnsureCurrentSemester(
	tahunAjaran *models.TahunAjaran,
) (*models.Semester, error) {

	now := time.Now()

	var (
		namaSemester string
		kode         string
	)

	if now.Month() >= time.July {

		namaSemester = "Ganjil"
		kode = "1"

	} else {

		namaSemester = "Genap"
		kode = "2"

	}

	var semester models.Semester

	err := s.DB.
		Where("tahun_ajaran_id = ?", tahunAjaran.ID).
		Where("kode = ?", kode).
		First(&semester).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {

		// Matikan semester lama
		if err := s.DB.
			Model(&models.Semester{}).
			Where("tahun_ajaran_id = ?", tahunAjaran.ID).
			Update("is_active", false).Error; err != nil {

			return nil, err
		}

		semester = models.Semester{
			TahunAjaranID: tahunAjaran.ID,
			NamaSemester:  namaSemester,
			Kode:          kode,
			IsActive:      true,
		}

		if err := s.DB.Create(&semester).Error; err != nil {
			return nil, fmt.Errorf(
				"gagal membuat semester: %w",
				err,
			)
		}

		return &semester, nil
	}

	if err != nil {
		return nil, err
	}

	// Pastikan semester aktif
	if err := s.DB.
		Model(&models.Semester{}).
		Where("tahun_ajaran_id = ?", tahunAjaran.ID).
		Where("id <> ?", semester.ID).
		Update("is_active", false).Error; err != nil {

		return nil, err
	}

	if err := s.DB.
		Model(&semester).
		Update("is_active", true).Error; err != nil {

		return nil, err
	}

	semester.IsActive = true

	return &semester, nil
}
