package repository

import (
	"github.com/alekssro/banking/banking/domain/entity"
	"github.com/alekssro/banking/banking/shared/errs"
)

// CustomerRepository interface
type AccountRepository interface {
	Save(entity.Account) (*entity.Account, *errs.AppError)
	FindByID(string) (*entity.Account, *errs.AppError)
}
