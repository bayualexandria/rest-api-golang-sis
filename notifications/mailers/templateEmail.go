package mailers

import (
	"bytes"
	"html/template"
)

type WelcomeMail struct {
	To       string
	Name     string
	Message  string
	Password string
	Username string
	URL      string
	NameButton string
	NameSubject string
}

func (w *WelcomeMail) Build() (string, string, string) {

	tmpl, _ := template.ParseFiles("views/mail/email.html")

	var body bytes.Buffer
	tmpl.Execute(&body, w)

	return w.To, w.NameSubject, body.String()
}
