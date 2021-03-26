package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alekssro/banking/dto"
	"github.com/alekssro/banking/service"
	"github.com/gorilla/mux"
)

// TransactionHandler struct defines the Transaction handler
// which depends on the service.TransactionService
type TransactionHandler struct {
	service service.TransactionService
}

func (th *TransactionHandler) newTransaction(w http.ResponseWriter, r *http.Request) {

	// get mux variables
	vars := mux.Vars(r)
	customerID := vars["customer_id"]
	accountID := vars["account_id"]

	var request dto.NewTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerID
		request.AccountID = accountID

		Transaction, appErr := th.service.NewTransaction(request)
		if appErr != nil {
			writeResponse(w, appErr.Code, appErr.Message)
		} else {
			fmt.Println(Transaction)
			writeResponse(w, http.StatusCreated, Transaction)
		}
	}
}
