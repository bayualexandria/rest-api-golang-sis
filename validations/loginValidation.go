package validations

import (
	"github.com/go-playground/validator/v10"
)

type LoginValidation struct {
	Username string `json:"username" binding:"required,numeric"`
	Password string `json:"password" binding:"required"`
}

var customMessages = map[string]string{
	"Username.required": "Username atau NIP harus diisi",
	"Username.numeric":  "Yang anda masukan bukan username atau NIP",
	"Password.required": "Password harus diisi",
}

func TranslateError(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			tag := fieldError.Tag()
			key := fieldName + "." + tag
			if msg, exists := customMessages[key]; exists {
				errors[fieldName] = msg
			}

		}
	}
	return errors
}
