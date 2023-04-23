package service_errors

import (
	"fmt"
	"net/http"
)

type ServiceError struct {
	StatusCode int
	Message    string
}

func NewError(statusCode int, message string) error {
	return ServiceError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func (e ServiceError) Error() string {
	return fmt.Sprintf("%v", e.Message)
}

func EmailValidationError(err error) error {
	return NewError(http.StatusBadRequest, fmt.Sprintf("invalid email format: %v", err.Error()))
}

func GetTransactionsError(err error) error {
	return NewError(http.StatusInternalServerError, fmt.Sprintf("get transactions error: %v", err.Error()))
}

func CalculateSummaryError(err error) error {
	return NewError(http.StatusInternalServerError, fmt.Sprintf("calculate summary error: %v", err.Error()))
}

func SendSummaryByEmailError(err error) error {
	return NewError(http.StatusInternalServerError, fmt.Sprintf("send summary by email error: %v", err.Error()))
}
