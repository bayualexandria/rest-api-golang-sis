package services

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

// EnsureCurrentYear memastikan tahun ajaran yang sesuai
// dengan tanggal sekarang tersedia dan aktif.
//
// Tahun ajaran dimulai pada bulan Juli.
//
// Januari - Juni:
//   contoh: 2026/2027
//
// Juli - Desember:
//   contoh: 2026/2027
func (s *TahunAjaranService) EnsureCurrentYear() (*models.TahunAjaran, error) {
	now := time.Now()

	// =====================================================
	// TENTUKAN TAHUN AWAL TAHUN AJARAN
	// =====================================================

	year := now.Year()

	// Januari - Juni masih menggunakan
	// tahun ajaran yang dimulai tahun sebelumnya.
	if now.Month() < time.July {
		year--
	}

	namaTahun := fmt.Sprintf("%d/%d", year, year+1)

	// =====================================================
	// TANGGAL TAHUN AJARAN
	// =====================================================

	startDate := time.Date(
		year,
		time.July,
		1,
		0,
		0,
		0,
		0,
		now.Location(),
	)

	endDate := time.Date(
		year+1,
		time.June,
		30,
		23,
		59,
		59,
		999999999,
		now.Location(),
	)

	var tahunAjaran models.TahunAjaran

	// =====================================================
	// TRANSACTION
	// =====================================================

	err := s.DB.Transaction(func(tx *gorm.DB) error {

		// =================================================
		// 1. CARI TAHUN AJARAN
		// =================================================

		err := tx.
			Where("nama_tahun = ?", namaTahun).
			First(&tahunAjaran).Error

		// =================================================
		// 2. BELUM ADA → BUAT BARU
		// =================================================

		if errors.Is(err, gorm.ErrRecordNotFound) {

			// Matikan seluruh tahun ajaran lama
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
				NamaTahun:     namaTahun,
				TanggalMulai:  startDate,
				TanggalSelesai: endDate,
				IsActive:      true,
			}

			if err := tx.Create(&tahunAjaran).Error; err != nil {
				return fmt.Errorf(
					"gagal membuat tahun ajaran %s: %w",
					namaTahun,
					err,
				)
			}

			return nil
		}

		// =================================================
		// 3. ERROR DATABASE
		// =================================================

		if err != nil {
			return fmt.Errorf(
				"gagal mencari tahun ajaran %s: %w",
				namaTahun,
				err,
			)
		}

		// =================================================
		// 4. NONAKTIFKAN TAHUN AJARAN LAIN
		// =================================================

		if err := tx.
			Model(&models.TahunAjaran{}).
			Where("id <> ?", tahunAjaran.ID).
			Update("is_active", false).Error; err != nil {

			return fmt.Errorf(
				"gagal menonaktifkan tahun ajaran lainnya: %w",
				err,
			)
		}

		// =================================================
		// 5. AKTIFKAN TAHUN AJARAN SEKARANG
		// =================================================

		if err := tx.
			Model(&models.TahunAjaran{}).
			Where("id = ?", tahunAjaran.ID).
			Update("is_active", true).Error; err != nil {

			return fmt.Errorf(
				"gagal mengaktifkan tahun ajaran %s: %w",
				namaTahun,
				err,
			)
		}

		tahunAjaran.IsActive = true

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &tahunAjaran, nil
}