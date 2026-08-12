package services

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
// untuk siswa_kelas aktif yang belum melakukan absensi hari ini.
func (s *AbsensiService) GenerateAlpa() error {

	// =====================================================
	// Waktu dan tanggal hari ini
	// =====================================================

	now := time.Now()

	tanggalHariIni := time.Date(
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
			return fmt.Errorf("tahun ajaran aktif tidak ditemukan")
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
			return fmt.Errorf("semester aktif tidak ditemukan")
		}

		return fmt.Errorf(
			"gagal mengambil semester aktif: %w",
			err,
		)
	}

	// =====================================================
	// Cari status ALPA
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
			"gagal mengambil status ALPA: %w",
			err,
		)
	}

	// =====================================================
	// Cari siswa_kelas aktif yang BELUM memiliki
	// absensi hari ini
	// =====================================================

	var siswaKelas []models.SiswaKelas

	subQuery := s.DB.
		Model(&models.AbsensiSiswa{}).
		Select("1").
		Where(
			"absensi_siswa.siswa_kelas_id = siswa_kelas.id",
		).
		Where(
			"absensi_siswa.tanggal = ?",
			tanggalHariIni,
		)

	if err := s.DB.
		Model(&models.SiswaKelas{}).
		Where("siswa_kelas.tahun_ajaran_id = ?", tahunAjaran.ID).
		Where("siswa_kelas.status = ?", "aktif").
		Where("NOT EXISTS (?)", subQuery).
		Find(&siswaKelas).Error; err != nil {

		return fmt.Errorf(
			"gagal mencari siswa yang belum absensi: %w",
			err,
		)
	}

	// =====================================================
	// Jika semua siswa sudah melakukan absensi
	// =====================================================

	if len(siswaKelas) == 0 {
		return nil
	}

	// =====================================================
	// Buat ALPA
	// =====================================================

	for _, siswa := range siswaKelas {

		absensiAlpa := models.AbsensiSiswa{
			SiswaKelasID:      siswa.ID,
			SemesterID:        semester.ID,
			StatusKehadiranID: statusAlpa.ID,
			Tanggal:           tanggalHariIni,
			JamMasuk:          nil,
			JamKeluar:         nil,
		}

		if err := s.DB.Create(&absensiAlpa).Error; err != nil {

			return fmt.Errorf(
				"gagal membuat ALPA siswa_kelas_id=%d: %w",
				siswa.ID,
				err,
			)
		}
	}

	return nil
}