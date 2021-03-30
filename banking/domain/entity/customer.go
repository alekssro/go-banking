package entity

// Customer struct defines the customer
type Customer struct {
	ID          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) StatusAsText() string {
	var statusAsText string
	if c.Status == "0" {
		statusAsText = "inactive"
	} else if c.Status == "1" {
		statusAsText = "active"
	}
	return statusAsText
}
