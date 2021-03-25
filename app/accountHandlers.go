package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alekssro/banking/dto"
	"github.com/alekssro/banking/service"
	"github.com/gorilla/mux"
)

// AccountHandler struct defines the account handler
// which depends on the service.AccountService
type AccountHandler struct {
	service service.AccountService
}

func (ah *AccountHandler) createAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customer_id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerID
		account, appErr := ah.service.NewAccount(request)
		if appErr != nil {
			writeResponse(w, appErr.Code, appErr.Message)
		} else {
			fmt.Println(account)
			writeResponse(w, http.StatusCreated, account)
		}
	}
}
