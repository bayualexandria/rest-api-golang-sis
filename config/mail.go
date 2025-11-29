package config

import (
	"os"

	"gopkg.in/gomail.v2"
)

func EmailConfig() *gomail.Dialer {

	return gomail.NewDialer(os.Getenv("MAIL_HOST"), 587, os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"))
}
