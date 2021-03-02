package domain

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
	FindAll() ([]Customer, error)
}
