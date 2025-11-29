package validations

import ("github.com/go-playground/validator/v10")

type LoginValidation struct {
    Username string `json:"username" binding:"required,numeric"`
    Password string `json:"password" binding:"required"`
}

var customMessages = map[string]string{
    "Username.required": "Username atau NIP harus diisi",
    "Username.numeric":  "Yang anda masukan bukan username atau NIP",
    "Password.required": "Password harus diisi",
}

func TranslateError(err error) string {
    if errs, ok := err.(validator.ValidationErrors); ok {
        for _, e := range errs {
            key := e.Field() + "." + e.ActualTag()
            if msg, exists := customMessages[key]; exists {
                return msg
            }
        }
    }
    return err.Error()
}

