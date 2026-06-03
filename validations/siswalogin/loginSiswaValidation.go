package siswalogin

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

type LoginSiswaValidation struct {
	Username string `form:"username" binding:"required,numeric"`
	Password string `form:"password" binding:"required"`
}

var customMessagesLoginSiswa = map[string]string{
	"Username.required": "Username atau NIS harus diisi",
	"Username.numeric":  "Yang anda masukan bukan username atau NIS",
	"Password.required": "Password harus diisi",
}

func TranslateErrorLoginSiswa(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {

			fieldName := fieldError.Field()
			jsonKey := toSnakeCase(fieldName)

			tag := fieldError.Tag()
			key := fieldName + "." + tag

			if msg, exists := customMessagesLoginSiswa[key]; exists {
				errors[jsonKey] = msg
			} else {
				errors[jsonKey] = fieldError.Error()
			}
			
		}
	}

	return errors
}

func toSnakeCase(str string) string {
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
