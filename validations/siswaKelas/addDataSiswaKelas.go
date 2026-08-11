package siswakelas

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

type AddDataSiswaKelasRequest struct {
	SiswaId uint64 `form:"siswa_id" binding:"required"`
	KelasId uint64 `form:"kelas_id" binding:"required"`
}

var addDataSiswaKelasMessages = map[string]string{
	"SiswaId.required": "Siswa ID wajib diisi.",
	"KelasId.required": "Kelas ID wajib diisi.",
}

func TranslateAddDataSiswaKelasError(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			jsonKey := toSnakeCaseAddDataSiswaKelas(fieldName)
			tag := fieldError.Tag()
			key := fieldName + "." + tag
			if msg, exists := addDataSiswaKelasMessages[key]; exists {
				errors[jsonKey] = msg
			}
		}
	}
	return errors
}

func toSnakeCaseAddDataSiswaKelas(str string) string {
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
