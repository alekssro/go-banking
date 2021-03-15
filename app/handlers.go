package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/alekssro/banking/service"
)

// Customer struct to store customer data
type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

// CustomerHandler struct defines the customer handler
// which depends on the service.CustomerService
type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "Rachel", City: "Alk", Zipcode: "28200"},
	// 	{Name: "Aleks", City: "Sanse", Zipcode: "28110"},
	// }

	customers, _ := ch.service.GetAllCustomers()

	// send response as xml if specified in Request Header
	if r.Header.Get("Content-Type") == "application/xml" {
		// xml writer
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}

	// send response as xml if specified in Request Header
	if r.Header.Get("Content-Type") == "application/json" {
		// Json writer
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
