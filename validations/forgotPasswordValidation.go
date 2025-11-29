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
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			key := e.Field() + "." + e.ActualTag()
			if msg, exists := forgotPasswordMessages[key]; exists {
				return msg
			}
		}
	}
	return err.Error()
}


