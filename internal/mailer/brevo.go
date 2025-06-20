package mailer

import (
	"bytes"
	"errors"
	"html/template"

	gomail "gopkg.in/mail.v2"
)

type brevoClient struct {
	fromEmail string
	username  string
	password  string
}

func NewBrevoClient(username, password, fromEmail string) (brevoClient, error) {
	if username == "" {
		return brevoClient{}, errors.New("brevo mail username is required")
	}

	if password == "" {
		return brevoClient{}, errors.New("brevo mail password is required")
	}

	return brevoClient{
		fromEmail: fromEmail,
		username:  username,
		password:  password,
	}, nil
}

func (m brevoClient) Send(templateFile, username, email string, data any, isSandbox bool) (int, error) {
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return -1, err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return -1, err
	}

	body := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(body, "body", data)
	if err != nil {
		return -1, err
	}

	message := gomail.NewMessage()
	message.SetHeader("From", m.fromEmail)
	message.SetHeader("To", email)
	message.SetHeader("Subject", subject.String())

	message.AddAlternative("text/html", body.String())

	dialer := gomail.NewDialer("smtp-relay.brevo.com", 587, m.username, m.password)

	if err := dialer.DialAndSend(message); err != nil {
		return -1, err
	}

	return 200, nil
}
