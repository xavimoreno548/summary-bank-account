package repository

import (
	"github.com/xavimoreno548/summary-bank-account/internal/model"
	"src/gorm.io/gorm"
)

type TransactionRepository interface {
	GetAll() ([]model.Transaction, error)
}

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepo{
		db: db,
	}
}

func (t TransactionRepo) GetAll() ([]model.Transaction, error) {
	return []model.Transaction{}, nil
}
