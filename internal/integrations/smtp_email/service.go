package smtp_email

import (
	"fmt"
	"net/smtp"
)

type service struct {
	serviceData ServiceData
}

func NewEmailService(serviceData ServiceData) IEmailService {
	return service{
		serviceData: serviceData,
	}
}

func (s service) BuildMessage(email string, htmlBody string, subject string) Message {
	from := s.serviceData.userName
	to := []string{email}
	message := fmt.Sprintf("To: %s\r\n", email)
	message += fmt.Sprintf("Subject: %s\r\n", subject)
	message += "MIME-version: 1.0;\r\n"
	message += "Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n"
	message += htmlBody

	return Message{
		From:    from,
		To:      to,
		Message: []byte(message),
	}
}

func (s service) Send(email Message) error {
	err := smtp.SendMail(s.serviceData.address, s.serviceData.auth, email.From, email.To, email.Message)
	if err != nil {
		return err
	}
	return nil
}
