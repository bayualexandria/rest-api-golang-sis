package validations

import (
	"github.com/go-playground/validator/v10"
)

type ForgotPasswordValidation struct {
	Email string `json:"email" binding:"required,email"`
}

var forgotPasswordMessages = map[string]string{
	"Email.required": "Email harus diisi",
	"Email.email":    "Format email tidak valid",
}

func TranslateForgotPasswordError(err error) string {
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
	return errors["Email"]
}


