package mapelController

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

type UpdateMapelValidation struct {
	Nama      string `form:"nama" binding:"omitempty"`
	Kode      string `form:"kode" binding:"omitempty"`
	Deskripsi string `form:"deskripsi" binding:"omitempty"`
}

var updateMapelMessages = map[string]string{
	"Nama.required":      "Nama Mapel wajib diisi.",
	"Kode.required":      "Kode Mapel wajib diisi.",
	"Deskripsi.required": "Deskripsi wajib diisi.",
}

func TranslateUpdateMapelError(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			jsonKey := toSnakeCaseUpdateMapel(fieldName)
			tag := fieldError.Tag()
			key := fieldName + "." + tag
			if msg, exists := updateMapelMessages[key]; exists {
				errors[jsonKey] = msg
			}
		}
	}
	return errors
}

func toSnakeCaseUpdateMapel(str string) string {
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
