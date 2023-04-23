package smtp_email

import "net/smtp"

type IEmailService interface {
	BuildMessage(email string, htmlBody string, subject string) Message
	Send(email Message) error
}

type Message struct {
	From    string
	To      []string
	Message []byte
}

type ServiceData struct {
	auth     smtp.Auth
	userName string
	address  string
}
