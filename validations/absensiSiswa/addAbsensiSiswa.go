package absensisiswa

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

type AddAbsensiValidation struct {
	Nis               string     `form:"nis" binding:"required"`
	Keterangan        *string `form:"keterangan"`
}

var addAbsensiMessages = map[string]string{
	"Nis.required":               "Siswa tidak diketahui.",
	"Keterangan.required":        "Keterangan wajib diisi.",
}

func TranslateAddAbsensiError(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			jsonKey := toSnakeCaseAddAbsensi(fieldName)
			tag := fieldError.Tag()
			key := fieldName + "." + tag
			if msg, exists := addAbsensiMessages[key]; exists {
				errors[jsonKey] = msg
			}
		}
	}
	return errors
}

func toSnakeCaseAddAbsensi(str string) string {
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
