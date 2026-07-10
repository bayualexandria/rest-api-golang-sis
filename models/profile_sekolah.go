package models

type ProfileSekolah struct {
	Id           int    `json:"id"`
	NamaSekolah  string `json:"nama_sekolah"`
	Alamat       string `json:"alamat"`
	NoTelp       string `json:"no_telp"`
	Akreditasi   string `json:"akreditasi"`
	ImageProfile string `json:"image_profile"`
	
}

func (ProfileSekolah) TableName() string {
	return "profile_sekolah" // jadi singular
}
