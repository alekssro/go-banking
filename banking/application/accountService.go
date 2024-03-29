package application

import (
	"github.com/alekssro/banking/banking/domain/repository"
	"github.com/alekssro/banking/lib/dto"
	"github.com/alekssro/banking/lib/errs"
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

// NewAccountService func adds a new default account service
func NewAccountService(repo repository.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
