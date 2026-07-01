package gurucontroller

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

type AddGuruValidation struct {
	Nip          string `form:"nip" binding:"required"`
	Nama         string `form:"nama" binding:"required"`
	JenisKelamin string `form:"jenis_kelamin" binding:"required,oneof=Laki-laki Perempuan"`
	NoHp         string `form:"no_hp" binding:"required,numeric"`
	Email        string `form:"email" binding:"required,email"`
	Alamat       string `form:"alamat" binding:"required"`
	StatusId     string `form:"status_id" binding:"required,oneof=1 2 3"` // 1=Admin, 2=Wali Kelas, 3=Guru
}

var addGuruMessages = map[string]string{
	"Nip.required":          "NIP wajib diisi.",
	"Nip.unique":            "NIP sudah digunakan.",
	"Nama.required":         "Nama wajib diisi.",
	"JenisKelamin.required": "Jenis kelamin wajib diisi.",
	"JenisKelamin.oneof":    "Jenis kelamin harus 'Laki-laki' atau 'Perempuan'.",
	"NoHp.required":         "No HP wajib diisi.",
	"NoHp.numeric":          "No HP harus berupa angka.",
	"Email.required":        "Email wajib diisi.",
	"Email.email":           "Format email tidak valid.",
	"Alamat.required":       "Alamat wajib diisi.",
	"StatusId.required":     "Status user wajib diisi.",
	"StatusId.oneof":        "Status user harus 1(Admin), 2(Wali Kelas), atau 3(Guru).",
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
