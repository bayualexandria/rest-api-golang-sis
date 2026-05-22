package notifications

import "backend-api/notifications/mailers"

func NotificationResetPassword(to, name, message, password, username string) error {
	mailer := mailers.NewMailer()
	email := mailers.WelcomeMail{
		To:          to,
		NameSubject: "Reset Password",
		Name:        name,
		Message:     message,
		Password:    password,
		Username:    username,
		URL:         "",
		NameButton:  "",
	}
	to, subject, body := email.Build()
	mailer.Send(to, subject, body)
	return nil

}
