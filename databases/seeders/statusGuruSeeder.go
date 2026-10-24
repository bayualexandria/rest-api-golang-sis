package seeders

import (
	"log"

	"gorm.io/gorm"
)

type StatusGuruSeeder struct {
	NamaStatus string
	Keterangan string
}

func (StatusGuruSeeder) TableName() string {
	return "status_guru" // jadi singular
}

func (s StatusGuruSeeder) Run(db *gorm.DB) {
	statusSiswaList := []StatusGuruSeeder{
		{NamaStatus: "Kepala Sekolah", Keterangan: "pimpinan lembaga pendidikan yang memegang kendali manajemen strategis sekolah"},
		{NamaStatus: "Wakil Kepala Sekolah", Keterangan: "Membantu kepala sekolah dalam bidang tertentu (seperti Kurikulum, Kesiswaan) atau memimpin program keahlian khusus (misalnya Ketua Jurusan TKJ, RPL, dll.)."},
		{NamaStatus: "Wali Kelas", Keterangan: "Wali Guru setiap kelas masing masing"},
		{NamaStatus: "Guru Tetap", Keterangan: "Memiliki status yang jelas dan resmi. Jika di sekolah negeri, berstatus sebagai ASN (PNS dengan pengangkatan tetap atau PPPK dengan sistem kontrak kerja). Jika di sekolah swasta, dikenal sebagai Guru Tetap Yayasan (GTY) dengan Surat Keputusan (SK) permanen"},
		{NamaStatus: "Guru Honorer", Keterangan: "Merupakan pegawai non-Aparatur Sipil Negara (non-ASN) atau non-PNS/PPPK yang diangkat oleh pejabat pembina kepegawaian atau kepala sekolah untuk jangka waktu tertentu."},
	}
	for _, status := range statusSiswaList {
		if err := db.Create(&status).Error; err != nil {
			log.Fatal("Error creating status guru:", err)
		}
	}
	log.Println("Seeder Status Guru selesai")
}
