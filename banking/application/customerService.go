package application

import (
	"github.com/alekssro/banking/banking/domain/repository"
	"github.com/alekssro/banking/banking/shared/dto"
	"github.com/alekssro/banking/banking/shared/errs"
)

// CustomerService defines a CustomerService interface
type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError)
	GetAllByStatus(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

// DefaultCustomerService struct defines the default customer service
// which depends on domain.CustomerRepository
type DefaultCustomerService struct {
	repo repository.CustomerRepository
}

// NewCustomerService func adds a new default customer service
func NewCustomerService(repo repository.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
