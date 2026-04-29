package mailers

import (
	"bytes"
	"html/template"
	"os"
)

type WelcomeMail struct {
	To          string
	Name        string
	Message     string
	Password    string
	Username    string
	URL         string
	NameButton  string
	NameSubject string
	NAMA_APP    string
}

func (w *WelcomeMail) Build() (string, string, string) {
	w.NAMA_APP = os.Getenv("NAMA_APP")

	tmpl, _ := template.ParseFiles("views/mail/email.html")

	var body bytes.Buffer

	tmpl.Execute(&body, w)

	return w.To, w.NameSubject, body.String()
}
