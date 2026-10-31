package seeders

import (
	"log"

	"gorm.io/gorm"
)

type JalurPendaftaranSeeder struct {
	NamaJalur string
	Deskripsi string
	Kuota     int
}

func (JalurPendaftaranSeeder) TableName() string {
	return "jalur_pendaftaran"
}

func (s JalurPendaftaranSeeder) Run(db *gorm.DB) {
	jalurList := []JalurPendaftaranSeeder{
		{
			NamaJalur: "Zonasi",
			Deskripsi: "Jalur pendaftaran berdasarkan radius jarak tempat tinggal calon peserta didik.",
			Kuota:     200,
		},
		{
			NamaJalur: "Afirmasi",
			Deskripsi: "Jalur pendaftaran bagi calon peserta didik dari keluarga kurang mampu.",
			Kuota:     75,
		},
		{
			NamaJalur: "Perpindahan Tugas",
			Deskripsi: "Jalur pendaftaran bagi anak yang mengikuti pindah tugas orang tua.",
			Kuota:     25,
		},
		{
			NamaJalur: "Prestasi",
			Deskripsi: "Jalur pendaftaran berdasarkan nilai rapor atau kejuaraan.",
			Kuota:     100,
		},
	}

	for _, jalur := range jalurList {
		// (Opsional) Cek dulu apakah data sudah ada untuk menghindari duplikasi saat seeding ulang
		var count int64
		db.Model(&JalurPendaftaranSeeder{}).Where("nama_jalur = ?", jalur.NamaJalur).Count(&count)
		
		if count == 0 {
			if err := db.Create(&jalur).Error; err != nil {
				log.Println("Gagal insert jalur pendaftaran:", err)
			}
		}
	}

	log.Println("Seeder JalurPendaftaran selesai dijalankan.")
}