package file_manager

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/xavimoreno548/summary-bank-account/internal/model"
)

const (
	CsvName = "data/txns.csv"
)

type IFileManager interface {
	GetRecordsFromFile(csvName string) ([]model.Transaction, error)
}

type service struct {
}

func NewFileManager() IFileManager {
	return &service{}
}

func (s service) GetRecordsFromFile(csvName string) ([]model.Transaction, error) {
	absPath, err := filepath.Abs(csvName)
	if err != nil {
		return []model.Transaction{}, err
	}

	file, err := os.Open(absPath)
	if err != nil {
		return []model.Transaction{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return []model.Transaction{}, err
	}

	var trxs []model.Transaction

	for _, record := range records[1:] {
		id, err := parseID(record[0])
		if err != nil {
			return []model.Transaction{}, err
		}

		date, err := parseDate(record[1])
		if err != nil {
			return []model.Transaction{}, err
		}

		transaction, err := parseTransaction(record[2])
		if err != nil {
			return []model.Transaction{}, err
		}

		t := model.Transaction{
			ID:          id,
			Date:        date,
			Transaction: transaction,
		}
		trxs = append(trxs, t)
	}

	return trxs, nil
}

func parseID(recordID string) (uint64, error) {
	id, err := strconv.ParseUint(recordID, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func parseDate(recordDate string) (time.Time, error) {
	recordDate = strings.Trim(recordDate, " ")
	timeLayout := "1/2/06"
	date, err := time.Parse(timeLayout, recordDate)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func parseTransaction(recordTrx string) (float64, error) {
	trxValue := strings.Trim(recordTrx, " ")

	if trxValue[0] == '+' {
		trxValue = trxValue[1:]
	}

	tValue, err := strconv.ParseFloat(trxValue, 64)
	if err != nil {
		return 0, err
	}

	return tValue, nil
}
