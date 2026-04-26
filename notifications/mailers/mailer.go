package mailers

import (
	"backend-api/config"

	"github.com/go-mail/mail"
)

type Mailer struct {
	Dialer *mail.Dialer
	From   string
}

func NewMailer() *Mailer {
	cfg := config.GetMailConfig()

	return &Mailer{
		Dialer: mail.NewDialer(
			cfg.Host,
			cfg.Port,
			cfg.Username,
			cfg.Password,
		),
		From: cfg.From,
	}
}

func (m *Mailer) Send(to, subject, body string) error {
	msg := mail.NewMessage()

	msg.SetHeader("From", msg.FormatAddress(m.From, "Administrator"))
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	return m.Dialer.DialAndSend(msg)
}
