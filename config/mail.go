package config

import ("gopkg.in/gomail.v2"
"os")

func mailer()  {
	m := gomail.NewMessage()

	m.SetHeader("From", os.Getenv("MAIL_EMAIL"))
	m.SetHeader("To", "wardanabayu455@gmail.com")
	m.SetHeader("Subject", "Test Email")
	m.SetBody("text/plain", "Halo ini email dari Golang")

	d := gomail.NewDialer(os.Getenv("MAIL_HOST"), 587, os.Getenv("MAIL_EMAIL"), os.Getenv("MAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}