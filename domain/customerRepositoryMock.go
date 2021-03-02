package domain

// CustomerRepositoryMock implements a CustomerRepository Mock connector
type CustomerRepositoryMock struct {
	customers []Customer
}

// FindAll method? implements the CustomerRepository interface for
// CustomerRepositoryMock struct
func (s CustomerRepositoryMock) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// NewCustomerRepositoryMock func implements adding Customers to the
// CustomerRepositoryMock
func NewCustomerRepositoryMock() CustomerRepositoryMock {
	customers := []Customer{
		{ID: "1", Name: "Alekss", City: "Sanse", Zipcode: "28700", DateofBirth: "1994-04-02", Status: "1"},
		{ID: "2", Name: "Reich", City: "Alk", Zipcode: "28200", DateofBirth: "1994-05-30", Status: "1"},
	}
	return CustomerRepositoryMock{customers: customers}
}
