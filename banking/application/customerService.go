package application

import (
	"github.com/alekssro/banking/banking/domain/entity"
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
	repo entity.CustomerRepository // TODO: Cambiar a repo
}

// GetAllCustomers method implements the CustomerService interface for
// DefaultCustomerService struct
func (s DefaultCustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	cs, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := []dto.CustomerResponse{}
	for _, c := range cs {
		response = append(response, c.ToDTO())
	}

	return response, nil
}

// GetAllByStatus method implements the CustomerService interface for the
// DefaultCustomerService struct. It returns the customers which has a correspondant status
func (s DefaultCustomerService) GetAllByStatus(status string) ([]dto.CustomerResponse, *errs.AppError) {
	// Translate to Repository vocabulary
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	}

	cs, err := s.repo.FindByStatus(status)
	if err != nil {
		return nil, err
	}

	response := []dto.CustomerResponse{}
	for _, c := range cs {
		response = append(response, c.ToDTO())
	}

	return response, nil
}

// GetCustomer method implements the CustomerService insterface for the
// DefaultCustomerService struct
func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ByID(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDTO()

	return &response, nil
}

// NewCustomerService func adds a new default customer service
func NewCustomerService(repo entity.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
