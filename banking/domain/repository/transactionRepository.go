package repository

import (
	"github.com/alekssro/banking/banking/domain/entity"
	"github.com/alekssro/banking/banking/shared/errs"
)

// TransactionRepository interface
type TransactionRepository interface {
	Withdrawal(entity.Transaction) (*entity.Transaction, *errs.AppError)
	Deposit(entity.Transaction) (*entity.Transaction, *errs.AppError)
}
