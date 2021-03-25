package app

import (
	"fmt"
	"net/http"

	"github.com/alekssro/banking/service"
)

// AccountHandler struct defines the account handler
// which depends on the service.AccountService
type AccountHandler struct {
	service service.AccountService
}

func (ah *AccountHandler) createAccount(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	query := url.Query()

	// TODO: Implement account creation
	fmt.Println(query)

}
