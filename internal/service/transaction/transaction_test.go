package transaction

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xavimoreno548/summary-bank-account/internal/repository/transaction"
	"github.com/xavimoreno548/summary-bank-account/internal/service/service_errors"
	"github.com/xavimoreno548/summary-bank-account/internal/service/transaction/summary"
)

func TestTransactionService(t *testing.T) {

	email := "example@example.com"
	// For happy path
	mockEH := summary.MockEmailHandler{}
	mockRepo := transaction.MockTransactionRepository{}
	// For email validation error
	mockEmailValidError := summary.MockEmailValidationError{}
	// For repo parse error
	mockRepoParseError := transaction.MockTransactionRepositoryError{}
	// For send email error
	mockEHSendError := summary.MockEmailSendError{}

	testCases := []struct {
		name    string
		repo    transaction.ITransactionRepository
		eh      summary.ITransactionEmailHandler
		errWant error
	}{
		{
			name: "happy_path",
			repo: mockRepo,
			eh:   mockEH,
		},
		{
			name:    "email_validation_error",
			repo:    mockRepo,
			eh:      mockEmailValidError,
			errWant: service_errors.EmailValidationError(fmt.Errorf("invlid email format")),
		},
		{
			name:    "repo_parse_error",
			repo:    mockRepoParseError,
			eh:      mockEH,
			errWant: service_errors.GetTransactionsError(fmt.Errorf("parse error")),
		},
		{
			name:    "send_email_error",
			repo:    mockRepo,
			eh:      mockEHSendError,
			errWant: service_errors.SendSummaryByEmailError(fmt.Errorf("error when try to send email")),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			serv := NewTransactionService(tc.repo, tc.eh)
			err := serv.SendSummaryByEmail(email)
			if tc.errWant != nil {
				assert.Equal(t, tc.errWant, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
