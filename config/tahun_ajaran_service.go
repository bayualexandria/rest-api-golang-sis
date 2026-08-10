package config

import (
	"backend-api/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type TahunAjaranService struct {
	DB *gorm.DB
}

func NewTahunAjaranService(db *gorm.DB) *TahunAjaranService {
	return &TahunAjaranService{
		DB: db,
	}
}

// EnsureCurrentYear memastikan tahun ajaran saat ini
// sudah tersedia di database.
func (s *TahunAjaranService) EnsureCurrentYear() (*models.TahunAjaran, error) {
	now := time.Now()

	// Tahun ajaran dimulai bulan Juli
	year := now.Year()
	if now.Month() < time.July {
		year--
	}

	namaTahun := fmt.Sprintf("%d/%d", year, year+1)

	startDate := time.Date(
		year,
		time.July,
		1,
		0, 0, 0, 0,
		now.Location(),
	)

	endDate := time.Date(
		year+1,
		time.June,
		30,
		23, 59, 59,
		0,
		now.Location(),
	)

	var tahunAjaran models.TahunAjaran

	err := s.DB.Transaction(func(tx *gorm.DB) error {

		// =====================================================
		// 1. Cari tahun ajaran berdasarkan nama
		// =====================================================

		err := tx.
			Where("nama_tahun = ?", namaTahun).
			First(&tahunAjaran).Error

		// =====================================================
		// 2. Jika belum ada → buat tahun ajaran baru
		// =====================================================

		if errors.Is(err, gorm.ErrRecordNotFound) {

			// Matikan semua tahun ajaran lama
			if err := tx.
				Model(&models.TahunAjaran{}).
				Where("is_active = ?", true).
				Update("is_active", false).Error; err != nil {

				return fmt.Errorf(
					"gagal menonaktifkan tahun ajaran lama: %w",
					err,
				)
			}

			// Buat tahun ajaran baru
			tahunAjaran = models.TahunAjaran{
				NamaTahun:       namaTahun,
				TanggalMulai:    startDate,
				TanggalSelesai:  endDate,
				IsActive:        true,
			}

			if err := tx.Create(&tahunAjaran).Error; err != nil {
				return fmt.Errorf(
					"gagal membuat tahun ajaran baru: %w",
					err,
				)
			}

			return nil
		}

		// =====================================================
		// 3. Jika terjadi error database
		// =====================================================

		if err != nil {
			return fmt.Errorf(
				"gagal mencari tahun ajaran: %w",
				err,
			)
		}

		// =====================================================
		// 4. Tahun ajaran sudah ada
		//    Pastikan dia aktif
		// =====================================================

		if err := tx.
			Model(&models.TahunAjaran{}).
			Where("id <> ?", tahunAjaran.ID).
			Update("is_active", false).Error; err != nil {

			return fmt.Errorf(
				"gagal menonaktifkan tahun ajaran lainnya: %w",
				err,
			)
		}

		if !tahunAjaran.IsActive {

			if err := tx.
				Model(&tahunAjaran).
				Update("is_active", true).Error; err != nil {

				return fmt.Errorf(
					"gagal mengaktifkan tahun ajaran: %w",
					err,
				)
			}

			tahunAjaran.IsActive = true
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &tahunAjaran, nil
}