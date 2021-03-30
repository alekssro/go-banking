package repository

import (
	"github.com/alekssro/banking/banking/domain/entity"
	"github.com/alekssro/banking/banking/shared/errs"
)

// CustomerRepository interface
type CustomerRepository interface {
	FindAll() ([]entity.Customer, *errs.AppError)
	FindByStatus(string) ([]entity.Customer, *errs.AppError)
	ByID(string) (*entity.Customer, *errs.AppError)
}
