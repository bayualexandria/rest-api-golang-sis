package usercalonsiswa

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

type LoginUserCalonSiswaValidation struct {
	Username string `form:"username" binding:"required,numeric"`
	Password string `form:"password" binding:"required"`
}

var customMessagesLoginUserCalonSiswa = map[string]string{
	"Username.required": "Username harus diisi",
	"Username.numeric":  "Yang anda masukan bukan No. Registrasi",
	"Password.required": "Password harus diisi",
}

func TranslateErrorLoginUserCalonSiswa(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {

			fieldName := fieldError.Field()
			jsonKey := toSnakeCaseLoginUserCalonSiswa(fieldName)

			tag := fieldError.Tag()
			key := fieldName + "." + tag

			if msg, exists := customMessagesLoginUserCalonSiswa[key]; exists {
				errors[jsonKey] = msg
			} else {
				errors[jsonKey] = fieldError.Error()
			}

		}
	}

	return errors

}

func toSnakeCaseLoginUserCalonSiswa(str string) string {
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
