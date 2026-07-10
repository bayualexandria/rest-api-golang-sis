package siswacontroller

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

type AddSiswaValidation struct {
	Nis          string                `form:"nis" binding:"required,numeric"`
	Nama         string                `form:"nama" binding:"required"`
	JenisKelamin string                `form:"jenis_kelamin" binding:"required,oneof=Laki-laki Perempuan"`
	NoHp         string                `form:"no_hp" binding:"required,numeric"`
	Email        string                `form:"email" binding:"required,email"`
	Alamat       string                `form:"alamat" binding:"required"`
}

var addSiswaMessages = map[string]string{
	"Nis.required":          "NIS wajib diisi.",
	"Nis.numeric":           "NIS harus berupa angka.",
	"Nama.required":         "Nama wajib diisi.",
	"JenisKelamin.required": "Jenis kelamin wajib diisi.",
	"JenisKelamin.oneof":    "Jenis kelamin harus 'Laki-laki' atau 'Perempuan'.",
	"NoHp.required":         "No HP wajib diisi.",
	"NoHp.numeric":          "No HP harus berupa angka.",
	"Email.required":        "Email wajib diisi.",
	"Email.email":           "Format email tidak valid.",
	"Alamat.required":       "Alamat wajib diisi.",
}

func TranslateAddSiswaError(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			jsonKey := toSnakeCaseAddSiswa(fieldName)
			tag := fieldError.Tag()
			key := fieldName + "." + tag
			if msg, exists := addSiswaMessages[key]; exists {
				errors[jsonKey] = msg
			}
		}
	}
	return errors
}

func toSnakeCaseAddSiswa(str string) string {
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
