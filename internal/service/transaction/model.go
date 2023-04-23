package transaction

type ITransactionService interface {
	SendSummaryByEmail(email string) error
}
