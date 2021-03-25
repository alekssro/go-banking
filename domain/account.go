package domain

import "github.com/alekssro/banking/errs"

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

// CustomerRepository interface
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
