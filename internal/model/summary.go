package model

type Summary struct {
	TotalBalance         float64           `json:"total_balance"`
	NumberOfTransactions []NumTransactions `json:"number_of_transactions"`
	AverageDebitAmount   float64           `json:"average_debit_amount"`
	AverageCreditAmount  float64           `json:"average_credit_amount"`
}

type NumTransactions struct {
	Month    string `json:"month"`
	Quantity int    `json:"quantity"`
}
