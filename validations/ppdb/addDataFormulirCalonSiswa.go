package ppdb

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

type AddDataFormulirCalonSiswaRequest struct {
	NamaLengkap      string `form:"nama_lengkap" binding:"required"`
	IdUserCalonSiswa string `form:"id_user_calon_siswa" binding:"required"`
	Nik              string `form:"nik" binding:"required"`
	Nisn             string `form:"nisn" binding:"required"`
	JenisKelamin     string `form:"jenis_kelamin" binding:"required"`
	Agama            string `form:"agama" binding:"required"`
	Ttl              string `form:"ttl" binding:"required"`
	Alamat           string `form:"alamat" binding:"required"`
	SekolahAsal      string `form:"sekolah_asal" binding:"required"`
	TahunLulus       string `form:"tahun_lulus" binding:"required"`
	NikAyah          string `form:"nik_ayah" binding:"required"`
	NamaAyah         string `form:"nama_ayah" binding:"required"`
	PekerjaanAyah    string `form:"pekerjaan_ayah" binding:"required"`
	NikIbu           string `form:"nik_ibu" binding:"required"`
	NamaIbu          string `form:"nama_ibu" binding:"required"`
	PekerjaanIbu     string `form:"pekerjaan_ibu" binding:"required"`
	NomorTelepon     string `form:"nomor_telepon" binding:"required"`
	JenisDokumen     []string `form:"jenis_dokumen" binding:"required"`
	
	NamaWali         string `form:"nama_wali" binding:"omitempty"`
	PekerjaanWali    string `form:"pekerjaan_wali" binding:"omitempty"`
}

var addDataFormulirCalonSiswaMessages = map[string]string{
	"NamaLengkap.required":      "Nama lengkap wajib diisi.",
	"IdUserCalonSiswa.required": "ID user calon siswa wajib diisi.",
	"Nik.required":              "NIK wajib diisi.",
	"Nisn.required":             "NISN wajib diisi.",
	"JenisKelamin.required":     "Jenis kelamin wajib diisi.",
	"Agama.required":            "Agama wajib diisi.",
	"Ttl.required":              "Tempat tanggal lahir wajib diisi.",
	"Alamat.required":           "Alamat wajib diisi.",
	"SekolahAsal.required":      "Sekolah asal wajib diisi.",
	"TahunLulus.required":       "Tahun lulus wajib diisi.",
	"NikAyah.required":          "NIK ayah wajib diisi.",
	"NamaAyah.required":         "Nama ayah wajib diisi.",
	"PekerjaanAyah.required":    "Pekerjaan ayah wajib diisi.",
	"NikIbu.required":           "NIK ibu wajib diisi.",
	"NamaIbu.required":          "Nama ibu wajib diisi.",
	"PekerjaanIbu.required":     "Pekerjaan ibu wajib diisi.",
	"NomorTelepon.required":     "Nomor telepon wajib diisi.",
	"JenisDokumen.required":     "Jenis dokumen wajib diisi.",
	
}

func GetAddDataFormulirCalonSiswaValidationError(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			jsonKey := toSnakeCaseAddDataFormulirCalonSiswa(fieldName)
			tag := fieldError.Tag()
			key := fieldName + "." + tag
			if msg, exists := addDataFormulirCalonSiswaMessages[key]; exists {
				errors[jsonKey] = msg
			}
		}
	}
	return errors
}

func toSnakeCaseAddDataFormulirCalonSiswa(str string) string {
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
