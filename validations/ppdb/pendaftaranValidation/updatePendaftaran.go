package pendaftaranValidation

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

type UpdatePendaftaranValidation struct {
	StatusVerifikasi  string `form:"status_verifikasi" binding:"omitempty,oneof=Pending Diverifikasi Ditolak"`
	CatatanVerifikasi string `form:"catatan_verifikasi" binding:"omitempty"`
}

var updatePendaftaranMessages = map[string]string{
	"StatusVerifikasi.required":  "Nama wajib diisi.",
	"CatatanVerifikasi.required": "Jenis kelamin wajib diisi.",
}

// Nama fungsi disesuaikan dari Siswa menjadi Guru agar seragam
func TranslateUpdatePendaftaranError(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {

			fieldName := fieldError.Field()
			jsonKey := toSnakeUpdatePendaftaranCase(fieldName)

			tag := fieldError.Tag()
			key := fieldName + "." + tag

			if msg, exists := updatePendaftaranMessages[key]; exists {
				errors[jsonKey] = msg
			} else {
				errors[jsonKey] = fieldError.Error()
			}
		}
	}

	return errors
}

func toSnakeUpdatePendaftaranCase(str string) string {
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
