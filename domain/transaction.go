package domain

import (
	"github.com/alekssro/banking/dto"
	"github.com/alekssro/banking/errs"
)

// Transaction struct defines the domain transaction entity
type Transaction struct {
	TransactionID   string `db:"transaction_id"`
	AccountID       string `db:"account_id"`
	Amount          float64
	TransactionType string `db:"transaction_type"`
	TransactionDate string `db:"transaction_date"`
}

func (t Transaction) ToTransactionDTO() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		TransactionId: t.TransactionID,
		AccountAmount: 0, // FIXME: Get account amount
	}
}

// TransactionRepository interface
type NewTransactionRepository interface {
	Transaction(Transaction) (*Transaction, *errs.AppError)
}
