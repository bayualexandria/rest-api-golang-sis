package usercalonsiswa

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

type RegisterUserCalonSiswaValidation struct {
	NamaLengkap string `form:"nama_lengkap" binding:"required"`
	Email       string `form:"email" binding:"required,email"`
	NomorHp     string `form:"nomor_hp" binding:"required,numeric"`

	Password string `form:"password" binding:"required"`
}

var customMessagesRegisterUserCalonSiswa = map[string]string{
	"NamaLengkap.required": "Nama lengkap harus diisi",
	"Email.required":       "Email harus diisi",
	"Email.email":          "Format email tidak valid",
	"NomorHp.required":     "Nomor HP harus diisi",
	"NomorHp.numeric":      "Nomor HP harus berupa angka",
	"Password.required":    "Password harus diisi",
}

func TranslateErrorRegisterUserCalonSiswa(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {

			fieldName := fieldError.Field()
			jsonKey := toSnakeCaseRegisterUserCalonSiswa(fieldName)

			tag := fieldError.Tag()
			key := fieldName + "." + tag

			if msg, exists := customMessagesRegisterUserCalonSiswa[key]; exists {
				errors[jsonKey] = msg
			} else {
				errors[jsonKey] = fieldError.Error()
			}

		}
	}

	return errors
}

func toSnakeCaseRegisterUserCalonSiswa(str string) string {
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
