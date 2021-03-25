package dto

import (
	"strings"

	"github.com/alekssro/banking/errs"
	"github.com/alekssro/banking/logger"
)

// Account struct defines the account for API response
type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("Insuficient amount to create an account (at least 5000.00 needed)")
	}
	logger.Info(strings.ToLower(r.AccountType))
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("Invalid account type, should be 'checking' or 'saving'")
	}
	return nil
}
