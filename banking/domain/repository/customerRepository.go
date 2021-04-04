package repository

import (
	"github.com/alekssro/banking/banking/domain/entity"
	"github.com/alekssro/banking/lib/errs"
)

// CustomerRepository interface
type CustomerRepository interface {
	FindAll() ([]entity.Customer, *errs.AppError)
	FindByStatus(string) ([]entity.Customer, *errs.AppError)
	ByID(string) (*entity.Customer, *errs.AppError)
}
