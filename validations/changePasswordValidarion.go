package validations

import "github.com/go-playground/validator/v10"

type ChangePasswordValidation struct {
	Password string `form:"password" binding:"required,min=8"`
}

var changePasswordMessages = map[string]string{
	"Password.required": "Password harus diisi",
	"Password.min":      "Password minimal 8 karakter",
}

func TranslateChangePasswordError(err error) string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			tag := fieldError.Tag()
			key := fieldName + "." + tag
			if msg, exists := changePasswordMessages[key]; exists {
				errors[fieldName] = msg
			}

		}
	}
	return errors["Password"]
}
