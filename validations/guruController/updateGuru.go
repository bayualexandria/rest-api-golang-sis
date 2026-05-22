package controllers // Sesuaikan dengan nama package kamu, misal: controllers atau validation

import (
	"mime/multipart"
	"unicode"

	"github.com/go-playground/validator/v10"
)

type UpdateGuruValidation struct {
	// Menambahkan binding:"omitempty" agar tidak wajib diisi semua saat update
	Nama         string                `form:"nama" binding:"omitempty"`
	JenisKelamin string                `form:"jenis_kelamin" binding:"omitempty,oneof=Laki-laki Perempuan"`
	NoHp         string                `form:"no_hp" binding:"omitempty,numeric"`
	Alamat       string                `form:"alamat" binding:"omitempty"`
	ImageProfile *multipart.FileHeader `form:"image_profile" binding:"omitempty"`
}

var updateGuruMessages = map[string]string{
	"Nama.required":         "Nama wajib diisi.",
	"JenisKelamin.required": "Jenis kelamin wajib diisi.",
	"JenisKelamin.oneof":    "Jenis kelamin harus 'Laki-laki' atau 'Perempuan'.",
	"NoHp.required":         "No HP wajib diisi.",
	"NoHp.numeric":          "No HP harus berupa angka.",
	"Alamat.required":       "Alamat wajib diisi.",
	"ImageProfile.required": "Image profile wajib diunggah.",
}

// Nama fungsi disesuaikan dari Siswa menjadi Guru agar seragam
func TranslateUpdateGuruError(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {

			fieldName := fieldError.Field()
			jsonKey := toSnakeCase(fieldName)

			tag := fieldError.Tag()
			key := fieldName + "." + tag

			if msg, exists := updateGuruMessages[key]; exists {
				errors[jsonKey] = msg
			} else {
				errors[jsonKey] = fieldError.Error()
			}
		}
	}

	return errors
}

func toSnakeCase(str string) string {
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
