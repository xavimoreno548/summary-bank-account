package transaction

import (
	"fmt"

	"github.com/xavimoreno548/summary-bank-account/internal/model"
)

type ITransactionRepository interface {
	GetTransactions() ([]model.Transaction, error)
}

type MockTransactionRepository struct{}

func (m MockTransactionRepository) GetTransactions() ([]model.Transaction, error) {
	return []model.Transaction{}, nil
}

type MockTransactionRepositoryError struct{}

func (m MockTransactionRepositoryError) GetTransactions() ([]model.Transaction, error) {
	return []model.Transaction{}, fmt.Errorf("parse error")
}
