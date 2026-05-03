package seeders

import (
	"log"

	"gorm.io/gorm"
)

type KelasSeeder struct {
	ID        int
	NamaKelas string
	Jurusan   string
}

func (KelasSeeder) TableName() string {
	return "kelas" // jadi singular
}

func (k KelasSeeder) Run(db *gorm.DB) {
	kelasList := []KelasSeeder{
		{ID: 1, NamaKelas: "X", Jurusan: "Teknik Komputer Jaringan"},
		{ID: 2, NamaKelas: "X", Jurusan: "Pemasaran"},

		{ID: 3, NamaKelas: "X", Jurusan: "Administrasi Perkantoran"},
		{ID: 4, NamaKelas: "X", Jurusan: "Akuntansi"},
		{ID: 5, NamaKelas: "XI", Jurusan: "Teknik Komputer Jaringan"},
		{ID: 6, NamaKelas: "XI", Jurusan: "Pemasaran"},
		{ID: 7, NamaKelas: "XI", Jurusan: "Administrasi Perkantoran"},
		{ID: 8, NamaKelas: "XI", Jurusan: "Akuntansi"},
		{ID: 9, NamaKelas: "XII", Jurusan: "Teknik Komputer Jaringan"},
		{ID: 10, NamaKelas: "XII", Jurusan: "Pemasaran"},
		{ID: 11, NamaKelas: "XII", Jurusan: "Administrasi Perkantoran"},
		{ID: 12, NamaKelas: "XII", Jurusan: "Akuntansi"},
	}
for _, kelas := range kelasList {
		if err := db.Create(&kelas).Error; err != nil {
			log.Fatal("Error creating kelas:", err)
		}
	}

	log.Println("Seeder Kelas selesai")
}
