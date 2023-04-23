package summary

import (
	"fmt"
	"strconv"

	"github.com/xavimoreno548/summary-bank-account/internal/model"
)

func CalculateSummary(transactions []model.Transaction) (model.Summary, error) {
	var balance, debitAcu, creditAcu float64
	var debitCount, creditCount int
	numTrx := make(map[string]int)

	for _, trx := range transactions {
		// Calculate the balance
		balance += trx.Transaction

		// Calculate debit and credit amounts
		if trx.Transaction < 0 {
			debitCount++
			debitAcu += trx.Transaction
		} else {
			creditCount++
			creditAcu += trx.Transaction
		}
		// Count transactions by month
		month := trx.Date.Month().String()
		numTrx[month] += 1
	}

	// Format Balance
	balance, err := formatFloatValues(balance)
	if err != nil {
		return model.Summary{}, err
	}
	// Average debit
	avgDebit, err := formatFloatValues(debitAcu / float64(debitCount))
	if err != nil {
		return model.Summary{}, err
	}
	// Average credit
	avgCredit, err := formatFloatValues(creditAcu / float64(creditCount))
	if err != nil {
		return model.Summary{}, err
	}
	// Number of transactions
	numTrxs := groupNumTransactionsByMap(numTrx)

	return model.Summary{
		TotalBalance:         balance,
		NumberOfTransactions: numTrxs,
		AverageDebitAmount:   avgDebit,
		AverageCreditAmount:  avgCredit,
	}, nil

}

func formatFloatValues(value float64) (float64, error) {
	vFormat := fmt.Sprintf("%.2f", value)
	vParsed, err := strconv.ParseFloat(vFormat, 64)
	if err != nil {
		return 0, err
	}
	return vParsed, nil
}

func groupNumTransactionsByMap(mapTrx map[string]int) []model.NumTransactions {
	numTransactions := make([]model.NumTransactions, 0)

	for month, quant := range mapTrx {
		oneTrx := model.NumTransactions{
			Month:    month,
			Quantity: quant,
		}
		numTransactions = append(numTransactions, oneTrx)
	}
	return numTransactions
}
