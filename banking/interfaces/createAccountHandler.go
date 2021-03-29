package interfaces

import (
	"encoding/json"
	"net/http"

	"github.com/alekssro/banking/banking/application"
	"github.com/alekssro/banking/banking/domain/entity"
	"github.com/alekssro/banking/banking/shared/dto"
	"github.com/alekssro/banking/banking/shared/logger"
	"github.com/alekssro/banking/banking/shared/message"
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
		message.WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerID
		account, appErr := ah.service.CreateAccount(request)
		if appErr != nil {
			message.WriteResponse(w, appErr.Code, appErr.Message)
		} else {
			logger.Info("New created account: id=" + account.AccountId)
			message.WriteResponse(w, http.StatusCreated, account)
		}
	}
}

func presentNewAccountDTO(a entity.Account) dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}
