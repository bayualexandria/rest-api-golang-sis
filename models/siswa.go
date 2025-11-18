package models

// Model User merepresentasikan tabel "users" di database
type Siswa struct {
	ID           int    `json:"id"`
	Nis          int    `json:"nis"`
	Nama         string `json:"nama"`
	JenisKelamin string `json:"jenis_kelamin"`
	NoHp         string `json:"no_hp"`
	Alamat       string `json:"alamat"`
	ImageProfile string `json:"image_profile"`
}

func (Siswa) TableName() string {
	return "siswa" // jadi singular
}
