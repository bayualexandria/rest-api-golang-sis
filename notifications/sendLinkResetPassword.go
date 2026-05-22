package notifications

import "backend-api/notifications/mailers"

func SendLinnkResetPassword(to, name, message, link, nameButton string) error {
	mailer := mailers.NewMailer()
	email := mailers.WelcomeMail{
		To:          to,
		NameSubject: "Forgot Password",
		Name:        name,
		Message:     message,
		Password:    "",
		Username:    "",
		URL:         link,
		NameButton:  nameButton,
	}
	to, subject, body := email.Build()
	mailer.Send(to, subject, body)
	return nil
}
