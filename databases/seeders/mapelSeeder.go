package seeders

import (
	"log"

	"gorm.io/gorm"
)

type MapelSeeder struct {
	Nama      string
	Kode      string
	Deskripsi string
}

func (MapelSeeder) TableName() string {
	return "mata_pelajaran"
}

func (k MapelSeeder) Run(db *gorm.DB) {
	mapelList := []MapelSeeder{
		// --------------------------------------------------------
		// 1. MATA PELAJARAN UMUM
		// --------------------------------------------------------
		{
			Nama:      "Pendidikan Agama dan Budi Pekerti",
			Kode:      "UMUM-PAI-01",
			Deskripsi: "Pembentukan karakter, moral, dan pemahaman nilai-nilai keagamaan.",
		},
		{
			Nama:      "Pendidikan Pancasila dan Kewarganegaraan",
			Kode:      "UMUM-PPKN-01",
			Deskripsi: "Pemahaman tata negara, hak dan kewajiban warga negara, serta nilai Pancasila.",
		},
		{
			Nama:      "Bahasa Indonesia",
			Kode:      "UMUM-BIN-01",
			Deskripsi: "Pengembangan tata bahasa, literasi, pemuatan teks teknis, dan komunikasi tertulis.",
		},
		{
			Nama:      "Matematika",
			Kode:      "UMUM-MTK-01",
			Deskripsi: "Matematika terapan, logika hitung, aljabar, dan statistik dasar.",
		},
		{
			Nama:      "Bahasa Inggris",
			Kode:      "UMUM-BIG-01",
			Deskripsi: "Penguasaan Bahasa Inggris umum dan komunikasi teknis dunia kerja (English for Business/IT).",
		},
		{
			Nama:      "Pendidikan Jasmani, Olahraga, dan Kesehatan",
			Kode:      "UMUM-PJOK-01",
			Deskripsi: "Pengembangan kebugaran jasmani, kesehatan, dan olahraga.",
		},
		{
			Nama:      "Sejarah Indonesia",
			Kode:      "UMUM-SEJ-01",
			Deskripsi: "Pemahaman sejarah perjuangan bangsa dan perkembangan nasional.",
		},
		{
			Nama:      "Seni Budaya",
			Kode:      "UMUM-SBK-01",
			Deskripsi: "Apresiasi karya seni, budaya lokal, dan kreativitas visual.",
		},

		// --------------------------------------------------------
		// 2. TEKNIK KOMPUTER DAN JARINGAN (TKJ)
		// --------------------------------------------------------
		{
			Nama:      "Dasar-Dasar Teknik Jaringan Komputer",
			Kode:      "TKJ-JARDAS-01",
			Deskripsi: "Pengenalan perangkat keras, kabel UTP/FO, model OSI, dan TCP/IP.",
		},
		{
			Nama:      "Administrasi Infrastruktur Jaringan",
			Kode:      "TKJ-AIJ-02",
			Deskripsi: "Konfigurasi VLAN, Routing (Static/Dynamic), NAT, Firewall, dan Bandwidth Management.",
		},
		{
			Nama:      "Administrasi Server Jaringan",
			Kode:      "TKJ-ASJ-03",
			Deskripsi: "Instalasi dan konfigurasi Server (Linux/Windows), Web Server, DNS, DHCP, FTP, dan Mail Server.",
		},
		{
			Nama:      "Teknologi Jaringan Berbasis Luas (WAN)",
			Kode:      "TKJ-TJBL-04",
			Deskripsi: "Infrastruktur jaringan WAN, jaringan kabel Fiber Optic, dan instalasi Nirkabel/Wireless.",
		},
		{
			Nama:      "Keamanan Jaringan Komputer",
			Kode:      "TKJ-SKJ-05",
			Deskripsi: "Prinsip kriptografi, penanganan ancaman siber, IDS/IPS, dan audit sistem keamanan.",
		},

		// --------------------------------------------------------
		// 3. PEMASARAN / BISNIS DARING DAN PEMASARAN (BDP)
		// --------------------------------------------------------
		{
			Nama:      "Dasar-Dasar Pemasaran",
			Kode:      "PM-PERDAS-01",
			Deskripsi: "Konsep dasar pemasaran, riset pasar, perilaku konsumen, dan segmentasi pasar.",
		},
		{
			Nama:      "Pemasaran Digital (Digital Marketing)",
			Kode:      "PM-DIGITAL-02",
			Deskripsi: "Pengelolaan SEO/SEM, Social Media Marketing, Content Creation, dan Meta/Google Ads.",
		},
		{
			Nama:      "Pengelolaan Bisnis Ritel",
			Kode:      "PM-RETAIL-03",
			Deskripsi: "Manajemen toko, tata letak barang (visual merchandising), stok barang, dan pelayanan pelanggan.",
		},
		{
			Nama:      "Strategi E-Commerce dan Marketplace",
			Kode:      "PM-ECOM-04",
			Deskripsi: "Manajemen toko online, strategi copywriting produk, dan optimalisasi platform e-commerce.",
		},
		{
			Nama:      "Penjualan dan Komunikasi Bisnis",
			Kode:      "PM-KOMBIS-05",
			Deskripsi: "Teknik negosiasi, presentasi produk, pelayanan prima (service excellence), dan penanganan komplain.",
		},

		// --------------------------------------------------------
		// 4. ADMINISTRASI PERKANTORAN / OTKP
		// --------------------------------------------------------
		{
			Nama:      "Dasar-Dasar Manajemen Perkantoran",
			Kode:      "AP-ADMUM-01",
			Deskripsi: "Konsep tata kelola kantor, manajemen dokumen, dan komunikasi perkantoran.",
		},
		{
			Nama:      "Pengelolaan Kearsipan",
			Kode:      "AP-ARSIP-02",
			Deskripsi: "Sistem penyimpanan, indeksasi, retensi dokumen fisik maupun digital (E-Archiving).",
		},
		{
			Nama:      "Otomatisasi Tata Kelola Kepegawaian",
			Kode:      "AP-KEPEG-03",
			Deskripsi: "Regulasi tenaga kerja, sistem penggajian (payroll), administrasi rekrutmen, dan penilaian kinerja.",
		},
		{
			Nama:      "Humas dan Keprotokolan",
			Kode:      "AP-HUMAS-04",
			Deskripsi: "Teknik Public Relations, penyelenggaraan acara (event organizer), dan etika keprotokolan.",
		},
		{
			Nama:      "Tata Kelola Sarana dan Prasarana",
			Kode:      "AP-SARPRAS-05",
			Deskripsi: "Perencanaan, pengadaan, inventarisasi, dan perawatan aset perkantoran.",
		},

		// --------------------------------------------------------
		// 5. AKUNTANSI DAN KEUANGAN (AKL)
		// --------------------------------------------------------
		{
			Nama:      "Akuntansi Dasar",
			Kode:      "AK-AKDAS-01",
			Deskripsi: "Persamaan dasar akuntansi, siklus akuntansi perusahaan jasa dan dagang.",
		},
		{
			Nama:      "Akuntansi Keuangan",
			Kode:      "AK-AKMAN-02",
			Deskripsi: "Pencatatan kas, piutang, persediaan barang, aset tetap, dan penyusunan laporan keuangan.",
		},
		{
			Nama:      "Komputer Akuntansi (MYOB/Accurate)",
			Kode:      "AK-MYOB-03",
			Deskripsi: "Pengoperasian software akuntansi untuk pengolahan transaksi dan pencetakan laporan otomatis.",
		},
		{
			Nama:      "Pajak dan Perpajakan",
			Kode:      "AK-PAJAK-04",
			Deskripsi: "Perhitungan PPh, PPN, pengisian SPT, dan regulasi perpajakan di Indonesia.",
		},
		{
			Nama:      "Akuntansi Lembaga / Instansi Pemerintah",
			Kode:      "AK-MANPUB-05",
			Deskripsi: "Sistem akuntansi keuangan daerah, pengelolaan anggaran publik, dan transparansi keuangan.",
		},
	}
	for _, kelas := range mapelList {
		if err := db.Create(&kelas).Error; err != nil {
			log.Fatal("Error creating kelas:", err)
		}
	}

	log.Println("Seeder Mapel selesai")
}
