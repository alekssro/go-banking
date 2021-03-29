package entity

import (
	"github.com/alekssro/banking/dto"
	"github.com/alekssro/banking/errs"
)

// Transaction struct defines the domain transaction entity
type Transaction struct {
	TransactionID   string `db:"transaction_id"`
	TransactionType string `db:"transaction_type"`
	TransactionDate string `db:"transaction_date"`
	Amount          float64
	AccountID       string `db:"account_id"`
	AccountBalance  float64
}

func (t Transaction) ToTransactionDTO() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		TransactionId:  t.TransactionID,
		AccountBalance: t.AccountBalance,
	}
}

// TransactionRepository interface
type NewTransactionRepository interface {
	Withdrawal(Transaction) (*Transaction, *errs.AppError)
	Deposit(Transaction) (*Transaction, *errs.AppError)
}
