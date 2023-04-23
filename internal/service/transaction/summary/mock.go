package summary

import (
	"fmt"

	"github.com/xavimoreno548/summary-bank-account/internal/model"
)

type MockEmailHandler struct{}

func (m MockEmailHandler) ValidateEmail(email string) error {
	return nil
}

func (m MockEmailHandler) LoadTemplate(templateName string, summary model.Summary) (string, error) {
	return "", nil
}

func (m MockEmailHandler) Send(email string, summary model.Summary) error {
	return nil
}

// MockEmailValidationError Mock for email validation error
type MockEmailValidationError struct{}

func (m MockEmailValidationError) ValidateEmail(email string) error {
	return fmt.Errorf("invlid email format")
}

func (m MockEmailValidationError) LoadTemplate(templateName string, summary model.Summary) (string, error) {
	return "", nil
}

func (m MockEmailValidationError) Send(email string, summary model.Summary) error {
	return nil
}

// MockEmailSendError Mock for send email error
type MockEmailSendError struct{}

func (m MockEmailSendError) ValidateEmail(email string) error {
	return nil
}

func (m MockEmailSendError) LoadTemplate(templateName string, summary model.Summary) (string, error) {
	return "", nil
}

func (m MockEmailSendError) Send(email string, summary model.Summary) error {
	return fmt.Errorf("error when try to send email")
}
