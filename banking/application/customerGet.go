package application

import (
	"github.com/alekssro/banking/banking/domain/entity"
	"github.com/alekssro/banking/lib/dto"
	"github.com/alekssro/banking/lib/errs"
)

// GetAllCustomers method implements the CustomerService interface for
// DefaultCustomerService struct
func (s DefaultCustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	cs, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := []dto.CustomerResponse{}
	for _, c := range cs {
		response = append(response, toCustomerResponseDTO(c))
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
		response = append(response, toCustomerResponseDTO(c))
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

	response := toCustomerResponseDTO(*c)

	return &response, nil
}

func toCustomerResponseDTO(c entity.Customer) dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.StatusAsText(),
	}
}
