package dto

import (
	"fmt"

	"github.com/alekssro/banking/banking/shared/errs"
	"github.com/alekssro/banking/banking/shared/logger"
)

// Transaction struct defines the Transaction for API response
type NewTransactionRequest struct {
	AccountID       string
	TransactionType string `json:"transaction_type"`
	TransactionDate string
	CustomerId      string
	Amount          float64 `json:"amount"`
}

func (r NewTransactionRequest) Validate() *errs.AppError {
	if r.Amount < 0 {
		logger.Error("Transaction amount is negative: " + fmt.Sprintf("%f", r.Amount))
		return errs.NewBadRequestError("Negative transaction amount")
	}
	if r.TransactionType != "withdrawal" && r.TransactionType != "deposit" {
		return errs.NewValidationError("Invalid transaction type: " + r.TransactionType)
	}

	return nil
}
