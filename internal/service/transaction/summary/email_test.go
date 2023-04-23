package summary

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xavimoreno548/summary-bank-account/internal/integrations/smtp_email"
)

func TestEmailValidate(t *testing.T) {

	email := "example@example.com"
	badEmail := "exampleexample.com"

	auth, err := smtp_email.NewGmailAuthSmtpFactory()
	assert.NoError(t, err)
	emailService := smtp_email.NewEmailService(auth)

	testCases := []struct {
		name    string
		es      smtp_email.IEmailService
		email   string
		errWant error
	}{
		{
			name:  "happy_path",
			es:    emailService,
			email: email,
		},
		{
			name:    "validation_error",
			es:      emailService,
			email:   badEmail,
			errWant: fmt.Errorf("mail: missing '@' or angle-addr"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			eh := NewEmailHandler(emailService)
			err = eh.ValidateEmail(tc.email)
			if tc.errWant != nil {
				assert.Equal(t, tc.errWant, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
