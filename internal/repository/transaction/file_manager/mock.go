package file_manager

import (
	"fmt"

	"github.com/xavimoreno548/summary-bank-account/internal/model"
)

type MockFileManager struct{}

func (m MockFileManager) GetRecordsFromFile(csvName string) ([]model.Transaction, error) {
	return []model.Transaction{}, nil
}

type MockFileManagerError struct{}

func (m MockFileManagerError) GetRecordsFromFile(csvName string) ([]model.Transaction, error) {
	return []model.Transaction{}, fmt.Errorf("error reading file")
}
