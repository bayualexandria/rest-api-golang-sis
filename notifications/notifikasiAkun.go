package notifications

import "backend-api/notifications/mailers"

func NotifikasiAkun(to, name, message string) error {
	mailer := mailers.NewMailer()

	email := mailers.WelcomeMail{
		To:          to,
		NameSubject: "Verifikasi Email",
		Name:        name,
		Message:     message,
		Password:    "",
		Username:    "",
		URL:         "https://www.smkxxxx.sch.id",
		NameButton:  "Kunjungi Website Kami",
	}

	to, subject, body := email.Build()
	mailer.Send(to, subject, body)
	return nil

}
