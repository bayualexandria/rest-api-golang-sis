package siswacontroller

import (
	"github.com/go-playground/validator/v10"
)

type UpdateSiswaValidation struct {
	Nama         string `json:"nama" binding:"required"`
	JenisKelamin string `json:"jenis_kelamin" binding:"required,oneof=Laki-laki Perempuan"`
	NoHp         string `json:"no_hp" binding:"required, numeric"`
	Alamat       string `json:"alamat" binding:"required"`
	ImageProfile string `json:"image_profile" binding:"required, image,mimes=jpeg,png,jpg|max=2048"`
	KelasID      int    `json:"kelas_id" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
}

var updateSiswaMessages = map[string]string{
	"Nama.required":         "Nama harus diisi",
	"JenisKelamin.required": "Jenis kelamin harus diisi",
	"JenisKelamin.oneof":    "Jenis kelamin harus Laki-laki atau Perempuan",
	"NoHp.required":         "No HP harus diisi",
	"NoHp.numeric":          "No HP harus berupa angka",
	"Alamat.required":       "Alamat harus diisi",
	"ImageProfile.required": "Image profile harus diisi",
	"ImageProfile.image":    "File harus berupa gambar",
	"ImageProfile.mimes":    "Format gambar harus jpeg, png, atau jpg",
	"ImageProfile.max":      "Ukuran gambar maksimal 2MB",
	"KelasID.required":      "Kelas ID harus diisi",
	"Email.required":        "Email harus diisi",
	"Email.email":           "Email harus valid",
}

func TranslateUpdateSiswaError(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			tag := fieldError.Tag()
			key := fieldName + "." + tag
			if msg, exists := updateSiswaMessages[key]; exists {
				errors[fieldName] = msg
			}
		}
	}
	return errors
}
