package dto

import (
	"fmt"

	"github.com/alekssro/banking/errs"
	"github.com/alekssro/banking/logger"
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
	// TODO: withdrawal amount should de available in the account
	//

	return nil
}
