package ruangkelas

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

type AddRuangKelasRequest struct {
	GuruWaliID    uint64 `json:"guru_wali_id" binding:"required"`
	KelasID       uint64 `json:"kelas_id" binding:"required"`
	TahunAjaranID uint64 `json:"tahun_ajaran_id" binding:"required"`
	SemesterID    uint64 `json:"semester_id" binding:"required"`
	Status        string `json:"status" binding:"required"`
}

var addRuangKelasMessages = map[string]string{
	"GuruWaliID.required":    "Guru Wali harus diisi",
	"KelasID.required":       "Kelas harus diisi",
	"TahunAjaranID.required": "Tahun Ajaran harus diisi",
	"SemesterID.required":    "Semester harus diisi",
	"Status.required":        "Status harus diisi",
}

func TranslateAddRuangKelasError(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			jsonKey := toSnakeCaseAddRuangKelas(fieldName)
			tag := fieldError.Tag()
			key := fieldName + "." + tag
			if msg, exists := addRuangKelasMessages[key]; exists {
				errors[jsonKey] = msg
			}
		}
	}
	return errors
}

func toSnakeCaseAddRuangKelas(str string) string {
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
