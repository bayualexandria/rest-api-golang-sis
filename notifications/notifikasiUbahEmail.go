package notifications

import "backend-api/notifications/mailers"

func NotificationUbahEmail(to, name, message string) error {
	mailer := mailers.NewMailer()
	email := mailers.WelcomeMail{
		To:          to,
		NameSubject: "Perubahan Email",
		Name:        name,
		Message:     message,
		Password:    "",
		Username:    "",
		URL:         "",
		NameButton:  "",
	}
	to, subject, body := email.Build()
	mailer.Send(to, subject, body)
	return nil

}
