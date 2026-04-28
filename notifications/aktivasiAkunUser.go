package notifications

import "backend-api/notifications/mailers"

func NotifikasiAktivasiAkunUser(to, name, message, url string) error {
	mailer := mailers.NewMailer()

	email := mailers.WelcomeMail{
		To:          to,
		NameSubject: "Aktivasi Akun",
		Name:        name,
		Message:     message,
		Password:    "",
		Username:    "",
		URL:         url,
		NameButton:  "Aktivasi Akun",
	}
	to, subject, body := email.Build()
	mailer.Send(to, subject, body)
	return nil
}
