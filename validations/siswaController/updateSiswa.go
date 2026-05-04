package siswacontroller

import (
	"mime/multipart"
	"unicode"

	"github.com/go-playground/validator/v10"
)

type UpdateSiswaValidation struct {
	Nama         string `form:"nama" binding:"required"`
	JenisKelamin string `form:"jenis_kelamin" binding:"required,oneof=Laki-laki Perempuan"`
	NoHp         string `form:"no_hp" binding:"required,numeric"`
	Alamat       string `form:"alamat" binding:"required"`
	Email        string `form:"email" binding:"required"`
	ImageProfile *multipart.FileHeader `form:"image_profile" binding:"required"`
}

var updateSiswaMessages = map[string]string{
	"Nama.required":         "Nama wajib diisi.",
	"JenisKelamin.required": "Jenis kelamin wajib diisi.",
	"JenisKelamin.oneof":    "Jenis kelamin harus 'Laki-laki' atau 'Perempuan'.",
	"NoHp.required":         "No HP wajib diisi.",
	"NoHp.numeric":          "No HP harus berupa angka.",
	"Alamat.required":       "Alamat wajib diisi.",
	"Email.required":        "Email harus diisi.",
	"ImageProfile.required": "Image profile wajib diunggah.",
}

func TranslateUpdateSiswaError(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {

			fieldName := fieldError.Field()
			jsonKey := toSnakeCase(fieldName)

			tag := fieldError.Tag()
			key := fieldName + "." + tag

			if msg, exists := updateSiswaMessages[key]; exists {
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
