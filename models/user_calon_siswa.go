package models

import (
	"time"
)

type UserCalonSiswa struct {
	NamaLengkap string `json:"nama_lengkap"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	NomorHp        string `json:"nomor_hp"`
	Password    string `json:"password"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (UserCalonSiswa) TableName() string {
	return "user_calon_siswa"
}
