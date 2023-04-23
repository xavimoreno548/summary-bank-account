package summary

import (
	"bytes"
	"html/template"
	"net/mail"
	"path/filepath"

	"github.com/xavimoreno548/summary-bank-account/internal/integrations/smtp_email"
	"github.com/xavimoreno548/summary-bank-account/internal/model"
)

const (
	Template = "internal/resources/summary.html"
	Subject  = "Summary Bank Account"
)

type ITransactionEmailHandler interface {
	ValidateEmail(email string) error
	LoadTemplate(templateName string, summary model.Summary) (string, error)
	Send(email string, summary model.Summary) error
}

type emailHandler struct {
	smtpClient smtp_email.IEmailService
}

func NewEmailHandler(smtpClient smtp_email.IEmailService) ITransactionEmailHandler {
	return &emailHandler{
		smtpClient: smtpClient,
	}
}

func (e emailHandler) ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}
	return nil
}

func (e emailHandler) LoadTemplate(templateName string, summary model.Summary) (string, error) {
	absPath, err := filepath.Abs(templateName)
	if err != nil {
		return "", err
	}

	tpl, err := template.ParseFiles(absPath)
	if err != nil {
		return "", err
	}

	var tplBuffer bytes.Buffer

	err = tpl.Execute(&tplBuffer, summary)
	if err != nil {
		return "", err
	}
	return tplBuffer.String(), nil
}

func (e emailHandler) Send(email string, summary model.Summary) error {
	htmlStringTemplate, err := e.LoadTemplate(Template, summary)
	if err != nil {
		return err
	}
	/*
		auth, err := smtp_email.NewGmailAuthSmtpFactory()
		if err != nil {
			return err
		}

		emailService := smtp_email.NewEmailService(auth)

	*/
	msg := e.smtpClient.BuildMessage(email, htmlStringTemplate, Subject)
	err = e.smtpClient.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
