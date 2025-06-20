package mailer

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
)

const (
	FromName            = "TMU-Webring"
	maxRetries          = 3
	UserWelcomeTemplate = "user_invitation.tmpl"
)

//go:embed "templates"
var FS embed.FS

type Client interface {
	Send(templateFile, username, email string, data any, isSandbox bool) (int, error)
}

func RenderEmailTemplate(name string, data any) (string, error) {
	tmpl, err := template.ParseFS(FS,
		"templates/base.html",
		fmt.Sprintf("templates/%s.html", name),
	)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	if err := tmpl.ExecuteTemplate(&out, "base", data); err != nil {
		return "", err
	}
	return out.String(), nil
}
