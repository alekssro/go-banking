package service

import (
	"github.com/alekssro/banking/domain"
	"github.com/alekssro/banking/errs"
)

// CustomerService defines a CustomerService interface
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

// DefaultCustomerService struct defines the default customer service
// which depends on domain.CustomerRepository
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomers method implements the CustomerService interface for
// DefaultCustomerService struct
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll()
}

// GetCustomer method implements the CustomerService insterface for the
// DefaultCustomerService struct
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ByID(id)
}

// NewCustomerService func adds a new default customer service
func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
