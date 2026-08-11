package siswakelas

import "github.com/go-playground/validator/v10"

type UpdateDataSiswaKelasRequest struct {
	SiswaId uint64 `form:"siswa_id" binding:"omitempty"`
	KelasId uint64 `form:"kelas_id" binding:"omitempty"`
}

var updateDataSiswaKelasMessages = map[string]string{
	"SiswaId.required": "Siswa ID wajib diisi.",
	"KelasId.required": "Kelas ID wajib diisi.",
}

func TranslateUpdateDataSiswaKelasError(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			jsonKey := toSnakeCaseUpdateDataSiswaKelas(fieldName)
			if msg, exists := updateDataSiswaKelasMessages[jsonKey+"."+fieldError.Tag()]; exists {
				errors[jsonKey] = msg
			} else {
				errors[jsonKey] = fieldError.Error()
			}
		}
	}

	return errors
}

func toSnakeCaseUpdateDataSiswaKelas(str string) string {
	var result []rune
	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return string(result)
}
