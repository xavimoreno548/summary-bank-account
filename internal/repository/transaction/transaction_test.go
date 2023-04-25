package transaction

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xavimoreno548/summary-bank-account/internal/model"
	"github.com/xavimoreno548/summary-bank-account/internal/repository/transaction/file_manager"
)

type MockFileManager struct{}

func (m MockFileManager) GetRecordsFromFile(filename string) ([]model.Transaction, error) {
	if filename == "data/txns.csv" {
		return nil, errors.New("error reading file")
	}
	return []model.Transaction{
		{
			ID:          1,
			Date:        time.Time{},
			Transaction: 0,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		},
	}, nil
}

func TestGetTransactions(t *testing.T) {

	testCases := []struct {
		name     string
		filename string
		fm       file_manager.IFileManager
		errWant  error
	}{
		{
			name:     "Success",
			filename: "test.csv",
			fm:       file_manager.MockFileManager{},
			errWant:  nil,
		},
		{
			name:     "FileError",
			filename: "data/txns.csv",
			fm:       file_manager.MockFileManagerError{},
			errWant:  errors.New("error reading file"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			trxRepo := NewTransactionRepository(nil, tc.fm)
			_, err := trxRepo.GetTransactions()
			if tc.errWant != nil {
				assert.Equal(t, tc.errWant, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
