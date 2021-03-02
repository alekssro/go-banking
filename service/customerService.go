package service

import "github.com/alekssro/banking/domain"

// CustomerService defines a CustomerService interface
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

// DefaultCustomerService struct defines the default customer service
// which depends on domain.CustomerRepository
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomers method? implements the CustomerService interface for
// DefaultCustomerService struct
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// NewCustomerService func adds a new default customer service
func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
