package ruangkelas

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

type UpdateRuangKelasRequest struct {
	GuruWaliId uint64 `form:"guru_wali_id" binding:"omitempty"`
	KelasId    uint64 `form:"kelas_id" binding:"omitempty"`
}

var updateRuangKelasMessages = map[string]string{
	"GuruWaliId.required": "Guru Wali harus diisi",
	"KelasId.required":    "Kelas harus diisi",
}

func TranslateUpdateRuangKelasError(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			jsonKey := toSnakeCaseUpdateRuangKelas(fieldName)
			tag := fieldError.Tag()
			key := fieldName + "." + tag
			if msg, exists := updateRuangKelasMessages[key]; exists {
				errors[jsonKey] = msg
			}
		}
	}
	return errors
}

func toSnakeCaseUpdateRuangKelas(str string) string {
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
