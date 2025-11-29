package validations

import (
	"github.com/go-playground/validator/v10"
)

type ResetPasswordValidation struct {
	Password string `json:"password" binding:"required,min=8"`
}

var resetPasswordMessages = map[string]string{
	"Password.required": "Password harus diisi",
	"Password.min":      "Password minimal 8 karakter",
}

func TranslateResetPasswordError(err error) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			key := e.Field() + "." + e.ActualTag()
			if msg, exists := resetPasswordMessages[key]; exists {
				return msg
			}
		}
	}
	return err.Error()
}
