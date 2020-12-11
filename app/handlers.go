package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// Customer struct to store customer data
type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Rachel", City: "Alk", Zipcode: "28200"},
		{Name: "Aleks", City: "Sanse", Zipcode: "28110"},
	}

	// send response as xml if specified in Request Header
	if r.Header.Get("Content-Type") == "application/xml" {
		// xml writer
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}

	// Json writer
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}
