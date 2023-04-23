package smtp_email

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/xavimoreno548/summary-bank-account/internal/integrations"
)

func NewGmailAuthSmtpFactory() (ServiceData, error) {
	err := integrations.LoadEnvVar()
	if err != nil {
		return ServiceData{}, err
	}

	userName := os.Getenv("GMAIL_USER_NAME")
	pass := os.Getenv("GMAIL_API_KEY")

	auth := smtp.PlainAuth("", userName, pass, "smtp.gmail.com")
	if auth == nil {
		return ServiceData{}, fmt.Errorf("error in gmail athentication")
	}

	address := os.Getenv("GMAIL_SMTP_ADDRESS")

	return ServiceData{
		auth:     auth,
		userName: userName,
		address:  address,
	}, nil
}
