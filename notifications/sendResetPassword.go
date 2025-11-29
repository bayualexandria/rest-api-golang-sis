package notifications

import (
	"backend-api/config"
	template "backend-api/templates"
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendResetPassword(to string, data interface{}) {

	mailer := config.EmailConfig()
	// Load HTML template
	body, err := template.LoadHTMLTemplate("./templates/mail/reset_password.html", data)
	if err != nil {
		return
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", os.Getenv("MAIL_USERNAME"))
	mail.SetHeader("Name", "Administrator")
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", "Forgot Password")

	mail.SetBody("text/html", body)
	if err := mailer.DialAndSend(mail); err != nil {
		fmt.Println("Failed to send email:", err)
		return
	}
}
