package api

import (
	"encoding/json"
	"net/http"

	"github.com/alekssro/banking/banking/application"
	"github.com/alekssro/banking/banking/shared/dto"
	"github.com/alekssro/banking/banking/shared/logger"
	"github.com/gorilla/mux"
)

// AccountHandler struct defines the account handler
// which depends on the application.AccountService
type AccountHandler struct {
	service application.AccountService
}

func (ah *AccountHandler) createAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customer_id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerID
		account, appErr := ah.service.CreateAccount(request)
		if appErr != nil {
			WriteResponse(w, appErr.Code, appErr.Message)
		} else {
			logger.Info("New created account: id=" + account.AccountId)
			WriteResponse(w, http.StatusCreated, account)
		}
	}
}
