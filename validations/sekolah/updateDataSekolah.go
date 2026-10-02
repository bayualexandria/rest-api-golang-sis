package sekolah

import (
	"mime/multipart"
	"unicode"

	"github.com/go-playground/validator/v10"
)

type UpdateDataSekolahValidation struct {
	NamaSekolah  string                `form:"nama_sekolah" binding:"omitempty"`
	Alamat       string                `form:"alamat" binding:"omitempty"`
	NoTelp       string                `form:"no_telp" binding:"omitempty"`
	Akreditasi   string                `form:"akreditasi" binding:"omitempty"`
	ImageProfile *multipart.FileHeader `form:"image_profile" binding:"omitempty"`
}

var UpdateDataSekolahValidationRules = map[string]string{
	"NamaSekolah.required":  "omitempty",
	"Alamat.required":       "omitempty",
	"NoTelp.required":       "omitempty",
	"Akreditasi.required":   "omitempty",
	"ImageProfile.required": "omitempty",
}

func TranslateUpdateDataSekolahError(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {

			fieldName := fieldError.Field()
			jsonKey := toSnakeUpdateSekolahCase(fieldName)

			tag := fieldError.Tag()
			key := fieldName + "." + tag

			if msg, exists := UpdateDataSekolahValidationRules[key]; exists {
				errors[jsonKey] = msg
			} else {
				errors[jsonKey] = fieldError.Error()
			}
		}
	}

	return errors
}

func toSnakeUpdateSekolahCase(str string) string {
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
