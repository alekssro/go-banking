package app

import (
	"encoding/json"
	"net/http"

	"github.com/alekssro/banking/service"
	"github.com/gorilla/mux"
)

// CustomerHandler struct defines the customer handler
// which depends on the service.CustomerService
type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		// Json writer
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		panic(err)
	}
}
