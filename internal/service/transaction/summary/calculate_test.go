package summary

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xavimoreno548/summary-bank-account/internal/model"
)

func TestCalculateSummary(t *testing.T) {

	transactions := []model.Transaction{
		{
			ID:          1,
			Date:        time.Date(2021, 5, 1, 0, 0, 0, 0, time.UTC),
			Transaction: 20,
		},
		{
			ID:          2,
			Date:        time.Date(2021, 5, 5, 0, 0, 0, 0, time.UTC),
			Transaction: -10.3,
		},
		{
			ID:          3,
			Date:        time.Date(2021, 6, 12, 0, 0, 0, 0, time.UTC),
			Transaction: -10.9,
		},
		{
			ID:          1,
			Date:        time.Date(2021, 7, 5, 0, 0, 0, 0, time.UTC),
			Transaction: 20,
		},
		{
			ID:          1,
			Date:        time.Date(2021, 8, 25, 0, 0, 0, 0, time.UTC),
			Transaction: -8.4,
		},
	}

	sumWant := model.Summary{
		TotalBalance: 10.40,
		NumberOfTransactions: []model.NumTransactions{
			{
				Month:    "May",
				Quantity: 2,
			}, {
				Month:    "June",
				Quantity: 1,
			}, {
				Month:    "July",
				Quantity: 1,
			}, {
				Month:    "August",
				Quantity: 1,
			},
		},
		AverageDebitAmount:  -9.87,
		AverageCreditAmount: 20,
	}

	summary, _ := CalculateSummary(transactions)
	assert.NotNil(t, summary)
	assert.Equal(t, sumWant, summary)
}
