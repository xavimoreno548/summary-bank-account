package transaction

import (
	"github.com/xavimoreno548/summary-bank-account/internal/repository/transaction"
	"github.com/xavimoreno548/summary-bank-account/internal/service/service_errors"
	"github.com/xavimoreno548/summary-bank-account/internal/service/transaction/summary"
)

type service struct {
	repository   transaction.ITransactionRepository
	emailHandler summary.ITransactionEmailHandler
}

func NewTransactionService(repository transaction.ITransactionRepository, emailHandler summary.ITransactionEmailHandler) ITransactionService {
	return &service{
		repository:   repository,
		emailHandler: emailHandler,
	}
}

func (s service) SendSummaryByEmail(email string) error {
	// Validate email
	err := s.emailHandler.ValidateEmail(email)
	if err != nil {
		return service_errors.EmailValidationError(err)
	}
	// Get records from repository
	transactions, err := s.repository.GetTransactions()
	if err != nil {
		return service_errors.GetTransactionsError(err)
	}
	// Calculate the summary
	accountSummary, err := summary.CalculateSummary(transactions)
	if err != nil {
		return service_errors.CalculateSummaryError(err)
	}
	// Send account summary in email
	err = s.emailHandler.Send(email, accountSummary)
	if err != nil {
		return service_errors.SendSummaryByEmailError(err)
	}

	return nil
}
