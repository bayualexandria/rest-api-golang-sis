package gurumapel

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

type AddGuruMapelRequest struct {
	GuruId          string `form:"guru_id" binding:"required"`
	MataPelajaranId string `form:"mata_pelajaran_id" binding:"required"`
	
}

var addGuruMapelMessages = map[string]string{
	"GuruId.required":        "Guru wajib diisi.",
	"MataPelajaranId.required":       "Mata Pelajaran wajib diisi.",
}

func TranslateAddGuruMapelError(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			jsonKey := toSnakeCaseAddGuruMapel(fieldName)
			tag := fieldError.Tag()
			key := fieldName + "." + tag
			if msg, exists := addGuruMapelMessages[key]; exists {
				errors[jsonKey] = msg
			}
		}
	}
	return errors
}

func toSnakeCaseAddGuruMapel(str string) string {
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
