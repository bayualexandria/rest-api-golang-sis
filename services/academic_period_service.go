package services

import (
	"backend-api/models"
	"fmt"

	"gorm.io/gorm"
)

type AcademicPeriodService struct {
	DB                 *gorm.DB
	TahunAjaranService *TahunAjaranService
	SemesterService    *SemesterService
}

func NewAcademicPeriodService(db *gorm.DB) *AcademicPeriodService {
	return &AcademicPeriodService{
		DB:                 db,
		TahunAjaranService: NewTahunAjaranService(db),
		SemesterService:    NewSemesterService(db),
	}
}

func (s *AcademicPeriodService) EnsureCurrentPeriod() (
	*models.TahunAjaran,
	*models.Semester,
	error,
) {

	tahunAjaran, err := s.TahunAjaranService.EnsureCurrentYear()
	if err != nil {
		return nil, nil, fmt.Errorf(
			"gagal memastikan tahun ajaran: %w",
			err,
		)
	}

	semester, err := s.SemesterService.EnsureCurrentSemester(tahunAjaran)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"gagal memastikan semester: %w",
			err,
		)
	}

	return tahunAjaran, semester, nil
}