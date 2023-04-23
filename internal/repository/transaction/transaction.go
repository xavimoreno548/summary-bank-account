package transaction

import (
	"github.com/xavimoreno548/summary-bank-account/internal/model"
	"github.com/xavimoreno548/summary-bank-account/internal/repository/transaction/file_manager"
	"gorm.io/gorm"
)

type trxRepo struct {
	db          *gorm.DB
	fileManager file_manager.IFileManager
}

func NewTransactionRepository(db *gorm.DB, fm file_manager.IFileManager) ITransactionRepository {
	return &trxRepo{
		db:          db,
		fileManager: fm,
	}
}

func (t trxRepo) GetTransactions() ([]model.Transaction, error) {
	// Get records form file
	transactions, err := t.fileManager.GetRecordsFromFile(file_manager.CsvName)
	if err != nil {
		return []model.Transaction{}, err
	}

	// Return records to service
	return transactions, nil
}
