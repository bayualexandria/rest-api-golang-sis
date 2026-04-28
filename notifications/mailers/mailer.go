package mailers

import (
	"backend-api/config"
	"crypto/tls"

	"github.com/go-mail/mail"
)

type Mailer struct {
	Dialer *mail.Dialer
	From   string
}

func NewMailer() *Mailer {
	cfg := config.GetMailConfig()

	dialer := mail.NewDialer(
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
	)

	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	return &Mailer{
		Dialer: dialer,
		From:   cfg.From,
	}
}

func (m *Mailer) Send(to, subject, body string) error {
	msg := mail.NewMessage()
	msg.SetHeader("From", msg.FormatAddress(m.From, "Administrator"))
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetHeader("MIME-Version", "1.0")
	msg.SetHeader("Content-Type", "text/html; charset=UTF-8")
	msg.SetBody("text/html", body)

	return m.Dialer.DialAndSend(msg)
}
