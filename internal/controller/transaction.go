package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xavimoreno548/summary-bank-account/internal/service/service_errors"
	"github.com/xavimoreno548/summary-bank-account/internal/service/transaction"
)

type ITransactionController interface {
	SendSummaryByEmail(ctx *gin.Context)
}

type transactionCtl struct {
	transactionService transaction.ITransactionService
}

func NewTransactionController(transactionService transaction.ITransactionService) ITransactionController {
	return &transactionCtl{
		transactionService: transactionService,
	}
}

func (t transactionCtl) SendSummaryByEmail(ctx *gin.Context) {
	// Send account summary by email
	email := ctx.Param("email")

	err := t.transactionService.SendSummaryByEmail(email)
	if err != nil {
		var serviceError service_errors.ServiceError
		if errors.As(err, &serviceError) {
			ctx.IndentedJSON(serviceError.StatusCode, gin.H{"error": serviceError.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	successMessage := fmt.Sprintf("succcess send summary to address %v", email)
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": successMessage})
	return
}
