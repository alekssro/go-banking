package domain

import "github.com/alekssro/banking/errs"

// Customer struct defines the customer
type Customer struct {
	ID          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

// CustomerRepository interface
type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindByStatus(string) ([]Customer, *errs.AppError)
	ByID(string) (*Customer, *errs.AppError)
}
