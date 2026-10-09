package jadwalpelajaran

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

type AddJadwalPelajaran struct {
	KelasId         uint64 `form:"kelas_id" binding:"required"`
	GuruId          uint64 `form:"guru_id" binding:"required"`
	MataPelajaranId uint64 `form:"mata_pelajaran_id" binding:"required"`
	Hari            string `form:"hari" binding:"required"`
	JamMulai        string `form:"jam_mulai" binding:"required"`
	JamSelesai      string `form:"jam_selesai" binding:"required"`
}

var addJadwalPelajaranMessages = map[string]string{
	"GuruId.required":          "Guru wajib diisi.",
	"MataPelajaranId.required": "Mata Pelajaran wajib diisi.",
	"KelasId.required":         "Kelas wajib diisi.",
	"Hari.required":            "Hari harus diisi",
	"JamMulai.required":        "Jam mulai wajib diisi.",
	"JamSelesai.required":      "Jam selesai harus diisi",
}

func TranslateAddJadwalPelajaranError(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			jsonKey := toSnakeCaseAddJadwalPelajaran(fieldName)
			tag := fieldError.Tag()
			key := fieldName + "." + tag
			if msg, exists := addJadwalPelajaranMessages[key]; exists {
				errors[jsonKey] = msg
			}
		}
	}
	return errors
}

func toSnakeCaseAddJadwalPelajaran(str string) string {
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
