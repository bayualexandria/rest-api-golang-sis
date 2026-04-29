package siswalogin

import (
	"github.com/go-playground/validator/v10"
)

type LoginSiswaValidation struct {
	Username string `json:"username" binding:"required,numeric"`
	Password string `json:"password" binding:"required"`
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
			tag := fieldError.Tag()
			key := fieldName + "." + tag
			if msg, exists := customMessagesLoginSiswa[key]; exists {
				errors[fieldName] = msg
			}

		}
	}
	return errors
}
