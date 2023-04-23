package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xavimoreno548/summary-bank-account/internal/controller"
	"github.com/xavimoreno548/summary-bank-account/internal/integrations/smtp_email"
	transactionRepo "github.com/xavimoreno548/summary-bank-account/internal/repository/transaction"
	"github.com/xavimoreno548/summary-bank-account/internal/repository/transaction/file_manager"
	"github.com/xavimoreno548/summary-bank-account/internal/service/transaction"
	"github.com/xavimoreno548/summary-bank-account/internal/service/transaction/summary"
)

func TransactionHandler() (controller.ITransactionController, error) {
	fileManager := file_manager.NewFileManager()
	transactionRepository := transactionRepo.NewTransactionRepository(nil, fileManager)

	emailServiceData, err := smtp_email.NewGmailAuthSmtpFactory()
	if err != nil {
		return nil, err
	}

	smtpClient := smtp_email.NewEmailService(emailServiceData)
	emailHandler := summary.NewEmailHandler(smtpClient)
	transactionService := transaction.NewTransactionService(transactionRepository, emailHandler)
	return controller.NewTransactionController(transactionService), nil
}

func AppRouter() (*gin.Engine, error) {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	trxController, err := TransactionHandler()
	if err != nil {
		return nil, err
	}
	router.POST("/send-summary/:email", trxController.SendSummaryByEmail)

	return router, nil
}

func RunApp(port string, router *gin.Engine) error {
	err := router.Run(":" + port)
	if err != nil {
		return fmt.Errorf("error when try to run server")
	}
	return nil
}
