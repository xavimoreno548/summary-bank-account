package service

import (
	"github.com/xavimoreno548/summary-bank-account/internal/model"
	"github.com/xavimoreno548/summary-bank-account/internal/repository"
)

type TransactionService interface {
	GetAll() ([]model.Transaction, error)
}

type service struct {
	repository repository.TransactionRepository
}

func NewTransactionService(repository repository.TransactionRepository) TransactionService {
	return &service{repository: repository}
}

func (s service) GetAll() ([]model.Transaction, error) {
	return []model.Transaction{}, nil
}
