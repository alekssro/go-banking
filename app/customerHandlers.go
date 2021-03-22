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

func (ch *CustomerHandler) queryCustomers(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	query := url.Query()
	status := query.Get("status")

	if status == "" {
		ch.getAllCustomers(w, r)
	} else {
		ch.getAllByStatus(w, r)
	}

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

func (ch *CustomerHandler) getAllByStatus(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllByStatus(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
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
