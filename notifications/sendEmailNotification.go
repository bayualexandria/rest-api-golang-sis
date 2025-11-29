package notifications

import (
	"backend-api/config"
	template "backend-api/templates"
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendVerificationEmail(to string, data interface{}) {

	mailer := config.EmailConfig()
	// Load HTML template
	body, err := template.LoadHTMLTemplate("./templates/mail/verify_email.html", data)
	if err != nil {
		return
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", os.Getenv("MAIL_USERNAME"))
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", "Email Verification")

	mail.SetBody("text/html", body)
	if err := mailer.DialAndSend(mail); err != nil {
		fmt.Println("Failed to send email:", err)
		return
	}
}
