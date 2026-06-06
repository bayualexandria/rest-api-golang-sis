package seeders

import (
	"log"

	"gorm.io/gorm"
)

type KelasSeeder struct {
	NamaKelas string
	Jurusan   string
}

func (KelasSeeder) TableName() string {
	return "kelas" // jadi singular
}

func (k KelasSeeder) Run(db *gorm.DB) {
	kelasList := []KelasSeeder{
		{NamaKelas: "X", Jurusan: "Teknik Komputer Jaringan"},
		{NamaKelas: "X", Jurusan: "Pemasaran"},

		{NamaKelas: "X", Jurusan: "Administrasi Perkantoran"},
		{NamaKelas: "X", Jurusan: "Akuntansi"},
		{NamaKelas: "XI", Jurusan: "Teknik Komputer Jaringan"},
		{NamaKelas: "XI", Jurusan: "Pemasaran"},
		{NamaKelas: "XI", Jurusan: "Administrasi Perkantoran"},
		{NamaKelas: "XI", Jurusan: "Akuntansi"},
		{NamaKelas: "XII", Jurusan: "Teknik Komputer Jaringan"},
		{NamaKelas: "XII", Jurusan: "Pemasaran"},
		{NamaKelas: "XII", Jurusan: "Administrasi Perkantoran"},
		{NamaKelas: "XII", Jurusan: "Akuntansi"},
	}
for _, kelas := range kelasList {
		if err := db.Create(&kelas).Error; err != nil {
			log.Fatal("Error creating kelas:", err)
		}
	}

	log.Println("Seeder Kelas selesai")
}
