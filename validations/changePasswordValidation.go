package validations

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

type ChangePasswordValidation struct {
	Password   string `form:"password" binding:"required,min=8"`
	RePassword string `form:"re_password" binding:"required,eqfield=Password"`
}

var changePasswordMessages = map[string]string{
	"Password.required":   "Password harus diisi",
	"Password.min":        "Password minimal 8 karakter",
	"RePassword.required": "Konfirmasi password harus diisi",
	"RePassword.eqfield":  "Konfirmasi password harus sama dengan password",
}

func TranslateChangePasswordError(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			jsonKey := toSnakeCaseChangePassword(fieldName)
			tag := fieldError.Tag()
			key := fieldName + "." + tag
			if msg, exists := changePasswordMessages[key]; exists {
				errors[jsonKey] = msg
			}
		}
	}
	return errors
}
func toSnakeCaseChangePassword(str string) string {
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
