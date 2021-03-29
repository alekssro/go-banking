package application

import (
	"github.com/alekssro/banking/banking/domain/repository"
	"github.com/alekssro/banking/banking/shared/dto"
	"github.com/alekssro/banking/banking/shared/errs"
)

// AccountService defines a AccountService interface
type AccountService interface {
	CreateAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	// FindAccount(dto.FindAccountRequest) (*dto.FindAccountResponse, *errs.AppError)
}

// DefaultAccountService struct defines the default Account service
// which depends on domain.AccountRepository
type DefaultAccountService struct {
	repo repository.AccountRepository
}
