
package config

import (
	"backend-api/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type AbsensiService struct {
	DB *gorm.DB
}

func NewAbsensiService(db *gorm.DB) *AbsensiService {
	return &AbsensiService{
		DB: db,
	}
}

// GenerateAlpa membuat absensi ALPA otomatis
// untuk siswa yang belum melakukan absensi
// setelah batas waktu absensi.
func (s *AbsensiService) GenerateAlpa() error {

	// =====================================================
	// Waktu sekarang
	// =====================================================

	now := time.Now()

	// =====================================================
	// BATAS WAKTU ABSENSI
	// Contoh: pukul 08:00
	// =====================================================

	jamBatas := 8

	if now.Hour() < jamBatas {

		// Belum waktunya membuat ALPA
		return nil
	}

	// =====================================================
	// Tanggal hari ini
	// =====================================================

	tanggal := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0,
		0,
		0,
		0,
		now.Location(),
	)

	// =====================================================
	// Cari tahun ajaran aktif
	// =====================================================

	var tahunAjaran models.TahunAjaran

	if err := s.DB.
		Where("is_active = ?", true).
		First(&tahunAjaran).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf(
				"tahun ajaran aktif tidak ditemukan",
			)
		}

		return fmt.Errorf(
			"gagal mengambil tahun ajaran aktif: %w",
			err,
		)
	}

	// =====================================================
	// Cari semester aktif
	// =====================================================

	var semester models.Semester

	if err := s.DB.
		Where("tahun_ajaran_id = ?", tahunAjaran.ID).
		Where("is_active = ?", true).
		First(&semester).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf(
				"semester aktif tidak ditemukan",
			)
		}

		return fmt.Errorf(
			"gagal mengambil semester aktif: %w",
			err,
		)
	}

	// =====================================================
	// Cari status kehadiran ALPA
	// =====================================================

	var statusAlpa models.StatusKehadiran

	if err := s.DB.
		Where("kode = ?", "A").
		First(&statusAlpa).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf(
				"status kehadiran ALPA dengan kode A tidak ditemukan",
			)
		}

		return fmt.Errorf(
			"gagal mengambil status kehadiran ALPA: %w",
			err,
		)
	}

	// =====================================================
	// Ambil semua siswa aktif
	// pada tahun ajaran aktif
	// =====================================================

	var siswaKelas []models.SiswaKelas

	if err := s.DB.
		Where("tahun_ajaran_id = ?", tahunAjaran.ID).
		Where("status = ?", "aktif").
		Find(&siswaKelas).Error; err != nil {

		return fmt.Errorf(
			"gagal mengambil siswa kelas aktif: %w",
			err,
		)
	}

	// =====================================================
	// Jika tidak ada siswa
	// =====================================================

	if len(siswaKelas) == 0 {
		return nil
	}

	// =====================================================
	// Proses setiap siswa
	// =====================================================

	for _, siswa := range siswaKelas {

		// -------------------------------------------------
		// Cek apakah sudah memiliki absensi hari ini
		// -------------------------------------------------

		var absensi models.AbsensiSiswa

		err := s.DB.
			Where("siswa_kelas_id = ?", siswa.ID).
			Where("tanggal = ?", tanggal).
			First(&absensi).Error

		// -------------------------------------------------
		// Sudah ada absensi
		// Jangan diubah menjadi ALPA
		// -------------------------------------------------

		if err == nil {
			continue
		}

		// -------------------------------------------------
		// Jika error bukan RecordNotFound
		// -------------------------------------------------

		if err != gorm.ErrRecordNotFound {

			return fmt.Errorf(
				"gagal mengecek absensi siswa_kelas_id=%d: %w",
				siswa.ID,
				err,
			)
		}

		// -------------------------------------------------
		// Belum ada absensi
		// Buat ALPA otomatis
		// -------------------------------------------------

		absensiAlpa := models.AbsensiSiswa{
			SiswaKelasID:      siswa.ID,
			SemesterID:        semester.ID,
			StatusKehadiranID: statusAlpa.ID,
			Tanggal:           tanggal,

			// ALPA tidak memiliki jam masuk
			JamMasuk: nil,

			// ALPA tidak memiliki jam keluar
			JamKeluar: nil,
		}

		if err := s.DB.Create(&absensiAlpa).Error; err != nil {

			return fmt.Errorf(
				"gagal membuat alpa siswa_kelas_id=%d: %w",
				siswa.ID,
				err,
			)
		}
	}

	return nil
}
