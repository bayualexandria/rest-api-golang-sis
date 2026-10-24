package gurucontroller

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

type AddGuruValidation struct {
	Nip          string  `form:"nip" binding:"required"`
	Nama         string  `form:"nama" binding:"required"`
	Ttl          string  `form:"ttl" binding:"required"`
	Pendidikan   string  `form:"pendidikan" binding:"required"`
	JenisKelamin *string `json:"jenis_kelamin" binding:"required,oneof=Laki-laki Perempuan"`
	NoHp         *string `json:"no_hp" binding:"required,numeric"`
	StatusId     *string `json:"status_id" binding:"required,oneof=1 2 3"`          // 1=Admin, 2=Wali Kelas, 3=Guru
	StatusGuruId *string `json:"status_guru_id" binding:"required,oneof=1 2 3 4 5"` //1=Kepala Sekolah, 2=Wakil Kepala Sekolah, 3=Guru Tetap, 4=Guru Honorer
	Email        string  `form:"email" binding:"required,email"`
	Alamat       string  `form:"alamat" binding:"required"`
}

var addGuruMessages = map[string]string{
	"Nip.required":          "NIP wajib diisi.",
	"Nip.unique":            "NIP sudah digunakan.",
	"Nama.required":         "Nama wajib diisi.",
	"Ttl.required":          "Tempat Tanggal Lahir wajib diisi.",
	"Pendidikan.required":   "Status Pendidikan wajib diisi.",
	"JenisKelamin.required": "Jenis kelamin wajib diisi.",
	"JenisKelamin.oneof":    "Jenis kelamin harus 'Laki-laki' atau 'Perempuan'.",
	"NoHp.required":         "No HP wajib diisi.",
	"NoHp.numeric":          "No HP harus berupa angka.",
	"Email.required":        "Email wajib diisi.",
	"Email.email":           "Format email tidak valid.",
	"Alamat.required":       "Alamat wajib diisi.",
	"StatusId.required":     "Status user wajib diisi.",
	"StatusId.oneof":        "Status user harus 1(Admin), 2(Wali Kelas), atau 3(Guru).",
	"StatusGuruId.required": "Status guru wajib diisi.",
	"StatusGuruId.oneof":    "Status guru harus 1(Kepala Sekolah), 2(Wakil Kepala Sekolah), 3(Wali Kelas), 4(Guru Tetap) dan 5(Guru Honorer).",
}

func TranslateAddGuruError(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			jsonKey := toSnakeCaseAddGuru(fieldName)
			tag := fieldError.Tag()
			key := fieldName + "." + tag
			if msg, exists := addGuruMessages[key]; exists {
				errors[jsonKey] = msg
			}
		}
	}
	return errors
}

func toSnakeCaseAddGuru(str string) string {
	var result []rune
	for i, r := range str {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
