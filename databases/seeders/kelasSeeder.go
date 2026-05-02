package seeders

import "gorm.io/gorm"

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
		{ID: 1, NamaKelas: "X IPA 1", Jurusan: "IPA"},
		{ID: 2, NamaKelas: "X IPA 2", Jurusan: "IPA"},

		{ID: 3, NamaKelas: "X IPS 1", Jurusan: "IPS"},
		{ID: 4, NamaKelas: "X IPS 2", Jurusan: "IPS"},
		{ID: 5, NamaKelas: "XI IPA 1", Jurusan: "IPA"},
		{ID: 6, NamaKelas: "XI IPA 2", Jurusan: "IPA"},
		{ID: 7, NamaKelas: "XI IPS 1", Jurusan: "IPS"},
		{ID: 8, NamaKelas: "XI IPS 2", Jurusan: "IPS"},
		{ID: 9, NamaKelas: "XII IPA 1", Jurusan: "IPA"},
		{ID: 10, NamaKelas: "XII IPA 2", Jurusan: "IPA"},
		{ID: 11, NamaKelas: "XII IPS 1", Jurusan: "IPS"},
		{ID: 12, NamaKelas: "XII IPS 2", Jurusan: "IPS"},
	}
	for _, kelas := range kelasList {
		db.Create(&kelas)
	}
}
