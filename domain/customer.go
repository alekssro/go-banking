package domain

import (
	"github.com/alekssro/banking/dto"
	"github.com/alekssro/banking/errs"
)

// Customer struct defines the customer
type Customer struct {
	ID          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusAsText() string {
	var statusAsText string
	if c.Status == "0" {
		statusAsText = "inactive"
	} else if c.Status == "1" {
		statusAsText = "active"
	}
	return statusAsText
}

func (c Customer) ToDTO() dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
}

// CustomerRepository interface
type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindByStatus(string) ([]Customer, *errs.AppError)
	ByID(string) (*Customer, *errs.AppError)
}
