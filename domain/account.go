package domain

import (
	"github.com/alekssro/banking/dto"
	"github.com/alekssro/banking/errs"
)

type Account struct {
	AccountId   string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount > amount
}

// CustomerRepository interface
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindByID(string) (*Account, *errs.AppError)
}
