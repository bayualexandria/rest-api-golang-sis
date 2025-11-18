package models

type Guru struct {
	ID           int    `json:"id"`
	Nip          int    `json:"nip"`
	Nama         string `json:"nama"`
	JenisKelamin string `json:"jenis_kelamin"`
	NoHp         string `json:"no_hp"`
	Alamat       string `json:"alamat"`
	ImageProfile string `json:"image_profile"`
}

func (Guru) TableName() string {
	return "guru" // jadi singular
}
