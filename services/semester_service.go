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

// EnsureCurrentSemester memastikan semester yang aktif
// sesuai dengan tanggal dan tahun ajaran saat ini.
func (s *SemesterService) EnsureCurrentSemester(
	tahunAjaran *models.TahunAjaran,
) (*models.Semester, error) {

	now := time.Now()

	var (
		namaSemester string
		kode         string
	)

	// Januari - Juni = Genap
	// Juli - Desember = Ganjil
	if now.Month() >= time.July {
		namaSemester = "Ganjil"
		kode = "1"
	} else {
		namaSemester = "Genap"
		kode = "2"
	}

	var semester models.Semester

	// Cari semester berdasarkan:
	// 1. Tahun ajaran
	// 2. Kode semester
	err := s.DB.
		Where("tahun_ajaran_id = ?", tahunAjaran.ID).
		Where("kode = ?", kode).
		First(&semester).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {

		// -------------------------------------------------
		// SEMESTER BELUM ADA
		// -------------------------------------------------

		// Nonaktifkan SEMUA semester yang aktif.
		//
		// Ini penting karena ketika tahun ajaran berganti,
		// semester tahun ajaran lama juga harus menjadi false.
		if err := s.DB.
			Model(&models.Semester{}).
			Where("is_active = ?", true).
			Update("is_active", false).Error; err != nil {

			return nil, fmt.Errorf(
				"gagal menonaktifkan semester lama: %w",
				err,
			)
		}

		// Buat semester baru
		semester = models.Semester{
			TahunAjaranID: tahunAjaran.ID,
			NamaSemester:  namaSemester,
			Kode:          kode,
			IsActive:      true,
		}

		if err := s.DB.Create(&semester).Error; err != nil {
			return nil, fmt.Errorf(
				"gagal membuat semester %s: %w",
				namaSemester,
				err,
			)
		}

		return &semester, nil
	}

	if err != nil {
		return nil, fmt.Errorf(
			"gagal mencari semester: %w",
			err,
		)
	}

	// -------------------------------------------------
	// SEMESTER SUDAH ADA
	// -------------------------------------------------

	// Nonaktifkan semua semester lain.
	//
	// Termasuk semester dari tahun ajaran sebelumnya.
	if err := s.DB.
		Model(&models.Semester{}).
		Where("id <> ?", semester.ID).
		Update("is_active", false).Error; err != nil {

		return nil, fmt.Errorf(
			"gagal menonaktifkan semester lainnya: %w",
			err,
		)
	}

	// Aktifkan semester yang sesuai dengan periode sekarang.
	if err := s.DB.
		Model(&models.Semester{}).
		Where("id = ?", semester.ID).
		Update("is_active", true).Error; err != nil {

		return nil, fmt.Errorf(
			"gagal mengaktifkan semester saat ini: %w",
			err,
		)
	}

	semester.IsActive = true

	return &semester, nil
}
