package siswacontroller

import (
	"mime/multipart"
	"unicode"

	"github.com/go-playground/validator/v10"
)

type UpdateSiswaValidation struct {
	Nama         string                `json:"nama" binding:"required"`
	JenisKelamin string                `json:"jenis_kelamin" binding:"required,oneof=Laki-laki Perempuan"`
	NoHp         string                `json:"no_hp" binding:"required,numeric"`
	Alamat       string                `json:"alamat" binding:"required"`
	ImageProfile *multipart.FileHeader `json:"image_profile" binding:"required"`
	KelasId      int                   `json:"kelas_id" binding:"required,gt=0"`
	Email        string                `json:"email" binding:"required,email"`
}

var updateSiswaMessages = map[string]string{
	"Nama.required":         "Nama wajib diisi.",
	"JenisKelamin.required": "Jenis kelamin wajib diisi.",
	"JenisKelamin.oneof":    "Jenis kelamin harus 'Laki-laki' atau 'Perempuan'.",
	"NoHp.required":         "No HP wajib diisi.",
	"NoHp.numeric":          "No HP harus berupa angka.",
	"Alamat.required":       "Alamat wajib diisi.",
	"ImageProfile.required": "Image profile wajib diunggah.",
	"KelasId.required":      "Kelas ID wajib diisi.",
	"KelasId.gt":            "Kelas ID harus lebih besar dari 0.",
	"Email.required":        "Email wajib diisi.",
	"Email.email":           "Format email tidak valid.",
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
