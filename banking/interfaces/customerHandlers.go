package interfaces

import (
	"net/http"

	"github.com/alekssro/banking/banking/application"
	"github.com/alekssro/banking/banking/shared/message"
	"github.com/gorilla/mux"
)

// CustomerHandler struct defines the customer handler
// which depends on the service.CustomerService
type CustomerHandler struct {
	service application.CustomerService
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
		message.WriteResponse(w, err.Code, err.AsMessage())
	} else {
		// Json writer
		message.WriteResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandler) getAllByStatus(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllByStatus(status)
	if err != nil {
		message.WriteResponse(w, err.Code, err.AsMessage())
	} else {
		message.WriteResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		message.WriteResponse(w, err.Code, err.AsMessage())
	} else {
		message.WriteResponse(w, http.StatusOK, customer)
	}
}
