package mapelController

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

type AddMapelValidation struct {
	Nama      string `form:"nama" binding:"required"`
	Kode      string `form:"kode" binding:"required"`
	Deskripsi string `form:"deskripsi" binding:"required"`
}

var addMapelMessages = map[string]string{
	"Nama.required":      "Nama Mapel wajib diisi.",
	"Kode.required":      "Kode Mapel wajib diisi.",
	"Deskripsi.required": "Deskripsi wajib diisi.",
}

func TranslateAddMapelError(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			jsonKey := toSnakeCaseAddMapel(fieldName)
			tag := fieldError.Tag()
			key := fieldName + "." + tag
			if msg, exists := addMapelMessages[key]; exists {
				errors[jsonKey] = msg
			}
		}
	}
	return errors
}

func toSnakeCaseAddMapel(str string) string {
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
