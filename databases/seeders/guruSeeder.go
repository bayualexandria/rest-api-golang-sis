package seeders

import (
	"backend-api/utils"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type GuruSeeder struct {
	Nip          string
	Nama         string
	JenisKelamin string
	NoHp         string
	Alamat       string
	ImageProfile string
}

type UserGuruSeeder struct {
	Name            string
	Username        string
	Email           string
	Password        string
	EmailVerifiedAt string
	StatusID        string
}

func (GuruSeeder) TableName() string {
	return "guru" // jadi singular
}

func (UserGuruSeeder) TableName() string {
	return "users" // jadi singular
}

func (s GuruSeeder) Run(db *gorm.DB) {
	passHash := "admin123" // Contoh password default

	// Daftar 20 data Guru
	guru := []GuruSeeder{
		{
			Nip:          "9106012508950001",
			Nama:         "Administrator",
			JenisKelamin: "Laki-laki",
			NoHp:         "081234567801",
			Alamat:       "Jl. Merdeka No. 10, Jakarta",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950002",
			Nama:         "Bagas Dwi Pratama",
			JenisKelamin: "Laki-laki",
			NoHp:         "081234567802",
			Alamat:       "Jl. Sudirman No. 15, Bandung",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950003",
			Nama:         "Sri Rahayu Ningsih",
			JenisKelamin: "Perempuan",
			NoHp:         "081234567803",
			Alamat:       "Jl. Ahmad Yani No. 20, Surabaya",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950004",
			Nama:         "Ahmad Fauzi",
			JenisKelamin: "Laki-laki",
			NoHp:         "081234567804",
			Alamat:       "Jl. Diponegoro No. 5, Yogyakarta",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950005",
			Nama:         "Siti Aminah",
			JenisKelamin: "Perempuan",
			NoHp:         "081234567805",
			Alamat:       "Jl. Pahlawan No. 12, Semarang",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950006",
			Nama:         "Eko Prasetyo",
			JenisKelamin: "Laki-laki",
			NoHp:         "081234567806",
			Alamat:       "Jl. Gajah Mada No. 8, Malang",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950007",
			Nama:         "Dewi Lestari",
			JenisKelamin: "Perempuan",
			NoHp:         "081234567807",
			Alamat:       "Jl. Hayam Wuruk No. 22, Solo",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950008",
			Nama:         "Rahmat Hidayat",
			JenisKelamin: "Laki-laki",
			NoHp:         "081234567808",
			Alamat:       "Jl. Imam Bonjol No. 14, Medan",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950009",
			Nama:         "Nurul Hidayah",
			JenisKelamin: "Perempuan",
			NoHp:         "081234567809",
			Alamat:       "Jl. Teuku Umar No. 9, Palembang",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950010",
			Nama:         "Hendra Setiawan",
			JenisKelamin: "Laki-laki",
			NoHp:         "081234567810",
			Alamat:       "Jl. Cut Nyak Dien No. 11, Banda Aceh",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950011",
			Nama:         "Siti Rahmawati",
			JenisKelamin: "Perempuan",
			NoHp:         "081234567811",
			Alamat:       "Jl. Veteran No. 33, Makassar",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950012",
			Nama:         "Joko Widodo",
			JenisKelamin: "Laki-laki",
			NoHp:         "081234567812",
			Alamat:       "Jl. Cendana No. 1, Surakarta",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950013",
			Nama:         "Rina Kartika",
			JenisKelamin: "Perempuan",
			NoHp:         "081234567813",
			Alamat:       "Jl. Melati No. 45, Denpasar",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950014",
			Nama:         "Budi Santoso",
			JenisKelamin: "Laki-laki",
			NoHp:         "081234567814",
			Alamat:       "Jl. Mawar No. 7, Bogor",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950015",
			Nama:         "Fitri Handayani",
			JenisKelamin: "Perempuan",
			NoHp:         "081234567815",
			Alamat:       "Jl. Kenanga No. 18, Bekasi",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950016",
			Nama:         "Agus Setiawan",
			JenisKelamin: "Laki-laki",
			NoHp:         "081234567816",
			Alamat:       "Jl. Anggrek No. 3, Tangerang",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950017",
			Nama:         "Siska Wulandari",
			JenisKelamin: "Perempuan",
			NoHp:         "081234567817",
			Alamat:       "Jl. Kamboja No. 25, Depok",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950018",
			Nama:         "Dedi Kurniawan",
			JenisKelamin: "Laki-laki",
			NoHp:         "081234567818",
			Alamat:       "Jl. Sisingamangaraja No. 12, Medan",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950019",
			Nama:         "Maya Sari",
			JenisKelamin: "Perempuan",
			NoHp:         "081234567819",
			Alamat:       "Jl. P Antasari No. 6, Banjarmasin",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950020",
			Nama:         "Rizky Ramadhan",
			JenisKelamin: "Laki-laki",
			NoHp:         "081234567820",
			Alamat:       "Jl. Soekarno Hatta No. 50, Bandung",
			ImageProfile: "storage/logo-pendidikan.png",
		},
	}

	// Generate 20 data User secara otomatis berdasarkan data guru di atas
	var user []UserGuruSeeder
	for i, g := range guru {
		// Menentukan status ID (misal berputar dari "1", "2", "3")
		statusID := fmt.Sprintf("%d", (i%3)+1)

		user = append(user, UserGuruSeeder{
			Name:            g.Nama,
			Username:        g.Nip,
			Email:           fmt.Sprintf("guru.penagihan%02d@gmail.com", i+1), // Contoh format email unik
			Password:        utils.HashPasswordUser(passHash),
			EmailVerifiedAt: time.Now().Format("2006-01-02"),
			StatusID:        statusID,
		})
	}

	if err := db.Create(&guru).Error; err != nil {
		log.Fatal("Error creating guru:", err)
	}
	if err := db.Create(&user).Error; err != nil {
		log.Fatal("Error creating user:", err)
	}

	log.Println("Seeder Guru selesai (20 data berhasil ditambahkan)")
}
