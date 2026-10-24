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
	Ttl          string
	Pendidikan   string
	NoHp         string
	Alamat       string
	ImageProfile string
	StatusGuruId string
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
			Ttl:          "Jakarta, 25 Agustus 1991",
			Pendidikan:   "S2 Manajemen Pendidikan",
			NoHp:         "081234567801",
			Alamat:       "Jl. Merdeka No. 10, Jakarta",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "1",
		},
		{
			Nip:          "9106012508950002",
			Nama:         "Bagas Dwi Pratama",
			JenisKelamin: "Laki-laki",
			Ttl:          "Bandung, 14 Februari 1992",
			Pendidikan:   "S1 Pendidikan Olahraga",
			NoHp:         "081234567802",
			Alamat:       "Jl. Sudirman No. 15, Bandung",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "2",
		},
		{
			Nip:          "9106012508950003",
			Nama:         "Sri Rahayu Ningsih",
			JenisKelamin: "Perempuan",
			Ttl:          "Surabaya, 10 Maret 1988",
			Pendidikan:   "S1 Pendidikan Bahasa Indonesia",
			NoHp:         "081234567803",
			Alamat:       "Jl. Ahmad Yani No. 20, Surabaya",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "3",
		},
		{
			Nip:          "9106012508950004",
			Nama:         "Ahmad Fauzi",
			JenisKelamin: "Laki-laki",
			Ttl:          "Yogyakarta, 5 Juli 1990",
			Pendidikan:   "S1 Pendidikan Matematika",
			NoHp:         "081234567804",
			Alamat:       "Jl. Diponegoro No. 5, Yogyakarta",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "4",
		},
		{
			Nip:          "9106012508950005",
			Nama:         "Siti Aminah",
			JenisKelamin: "Perempuan",
			Ttl:          "Semarang, 12 Desember 1989",
			Pendidikan:   "S1 Pendidikan Biologi",
			NoHp:         "081234567805",
			Alamat:       "Jl. Pahlawan No. 12, Semarang",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "3",
		},
		{
			Nip:          "9106012508950006",
			Nama:         "Eko Prasetyo",
			JenisKelamin: "Laki-laki",
			Ttl:          "Malang, 18 Januari 1991",
			Pendidikan:   "S1 Teknik Informatika",
			NoHp:         "081234567806",
			Alamat:       "Jl. Gajah Mada No. 8, Malang",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "3",
		},
		{
			Nip:          "9106012508950007",
			Nama:         "Dewi Lestari",
			JenisKelamin: "Perempuan",
			Ttl:          "Solo, 22 April 1993",
			Pendidikan:   "S1 Pendidikan Ekonomi",
			NoHp:         "081234567807",
			Alamat:       "Jl. Hayam Wuruk No. 22, Solo",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "3",
		},
		{
			Nip:          "9106012508950008",
			Nama:         "Rahmat Hidayat",
			JenisKelamin: "Laki-laki",
			Ttl:          "Medan, 30 Juni 1987",
			Pendidikan:   "S1 Pendidikan Sejarah",
			NoHp:         "081234567808",
			Alamat:       "Jl. Imam Bonjol No. 14, Medan",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "3",
		},
		{
			Nip:          "9106012508950009",
			Nama:         "Nurul Hidayah",
			JenisKelamin: "Perempuan",
			Ttl:          "Palembang, 8 September 1994",
			Pendidikan:   "S1 Pendidikan Agama Islam",
			NoHp:         "081234567809",
			Alamat:       "Jl. Teuku Umar No. 9, Palembang",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "4",
		},
		{
			Nip:          "9106012508950010",
			Nama:         "Hendra Setiawan",
			JenisKelamin: "Laki-laki",
			Ttl:          "Banda Aceh, 15 November 1989",
			Pendidikan:   "S1 Pendidikan Fisika",
			NoHp:         "081234567810",
			Alamat:       "Jl. Cut Nyak Dien No. 11, Banda Aceh",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "3",
		},
		{
			Nip:          "9106012508950011",
			Nama:         "Siti Rahmawati",
			JenisKelamin: "Perempuan",
			Ttl:          "Makassar, 2 Mei 1992",
			Pendidikan:   "S1 Pendidikan Kimia",
			NoHp:         "081234567811",
			Alamat:       "Jl. Veteran No. 33, Makassar",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "3",
		},
		{
			Nip:          "9106012508950012",
			Nama:         "Joko Widodo",
			JenisKelamin: "Laki-laki",
			Ttl:          "Surakarta, 21 Juni 1986",
			Pendidikan:   "S1 Kehutanan",
			NoHp:         "081234567812",
			Alamat:       "Jl. Cendana No. 1, Surakarta",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "3",
		},
		{
			Nip:          "9106012508950013",
			Nama:         "Rina Kartika",
			JenisKelamin: "Perempuan",
			Ttl:          "Denpasar, 9 Oktober 1995",
			Pendidikan:   "S1 Pendidikan Seni Bahasa",
			NoHp:         "081234567813",
			Alamat:       "Jl. Melati No. 45, Denpasar",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "3",
		},
		{
			Nip:          "9106012508950014",
			Nama:         "Budi Santoso",
			JenisKelamin: "Laki-laki",
			Ttl:          "Bogor, 17 Agustus 1990",
			Pendidikan:   "S1 Pendidikan Kewarganegaraan",
			NoHp:         "081234567814",
			Alamat:       "Jl. Mawar No. 7, Bogor",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "3",
		},
		{
			Nip:          "9106012508950015",
			Nama:         "Fitri Handayani",
			JenisKelamin: "Perempuan",
			Ttl:          "Bekasi, 25 Maret 1993",
			Pendidikan:   "S1 Bimbingan Konseling",
			NoHp:         "081234567815",
			Alamat:       "Jl. Kenanga No. 8, Bekasi",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "3",
		},
		{
			Nip:          "9106012508950016",
			Nama:         "Agus Setiawan",
			JenisKelamin: "Laki-laki",
			Ttl:          "Tangerang, 11 Juli 1991",
			Pendidikan:   "S1 Pendidikan Geografi",
			NoHp:         "081234567816",
			Alamat:       "Jl. Anggrek No. 3, Tangerang",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "3",
		},
		{
			Nip:          "9106012508950017",
			Nama:         "Siska Wulandari",
			JenisKelamin: "Perempuan",
			Ttl:          "Depok, 4 Januari 1994",
			Pendidikan:   "S1 Pendidikan Sastra Inggris",
			NoHp:         "081234567817",
			Alamat:       "Jl. Kamboja No. 25, Depok",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "3",
		},
		{
			Nip:          "9106012508950018",
			Nama:         "Dedi Kurniawan",
			JenisKelamin: "Laki-laki",
			Ttl:          "Medan, 29 September 1988",
			Pendidikan:   "S1 Pendidikan Teknik Mesin",
			NoHp:         "081234567818",
			Alamat:       "Jl. Sisingamangaraja No. 12, Medan",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "3",
		},
		{
			Nip:          "9106012508950019",
			Nama:         "Maya Sari",
			JenisKelamin: "Perempuan",
			Ttl:          "Banjarmasin, 13 November 1992",
			Pendidikan:   "S1 Pendidikan Guru PAUD",
			NoHp:         "081234567819",
			Alamat:       "Jl. P Antasari No. 6, Banjarmasin",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "3",
		},
		{
			Nip:          "9106012508950020",
			Nama:         "Rizky Ramadhan",
			JenisKelamin: "Laki-laki",
			Ttl:          "Bandung, 20 Mei 1995",
			Pendidikan:   "S1 Pendidikan Teknologi Informasi",
			NoHp:         "081234567820",
			Alamat:       "Jl. Soekarno Hatta No. 50, Bandung",
			ImageProfile: "storage/logo-pendidikan.png",
			StatusGuruId: "3",
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
