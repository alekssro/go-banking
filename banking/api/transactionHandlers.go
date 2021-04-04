package api

import (
	"encoding/json"
	"net/http"

	"github.com/alekssro/banking/banking/application"
	"github.com/alekssro/banking/lib/dto"
	"github.com/alekssro/banking/lib/logger"
	"github.com/gorilla/mux"
)

// TransactionHandler struct defines the Transaction handler
// which depends on the service.TransactionService
type TransactionHandler struct {
	service application.TransactionService
}

func (th *TransactionHandler) newTransaction(w http.ResponseWriter, r *http.Request) {

	// get mux variables
	vars := mux.Vars(r)
	customerID := vars["customer_id"]
	accountID := vars["account_id"]

	var request dto.NewTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerID
		request.AccountID = accountID
		Transaction, appErr := th.service.NewTransaction(request)
		if appErr != nil {
			WriteResponse(w, appErr.Code, appErr.Message)
		} else {
			logger.Info("New transaction: id=" + Transaction.TransactionId)
			WriteResponse(w, http.StatusCreated, Transaction)
		}
	}
}
