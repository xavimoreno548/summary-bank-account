package transaction

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xavimoreno548/summary-bank-account/internal/model"
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
		errWant  error
	}{
		{
			name:     "Success",
			filename: "test.csv",
			errWant:  nil,
		},
		{
			name:     "FileError",
			filename: "data/txns.csv",
			errWant:  errors.New("error reading file"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fm := MockFileManager{}
			trxRepo := NewTransactionRepository(nil, fm)
			transactions, err := trxRepo.GetTransactions()
			if tc.errWant != nil {
				assert.Equal(t, tc.errWant, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, transactions)
			}
		})
	}
}
