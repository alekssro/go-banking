package service

import (
	"github.com/alekssro/banking/domain"
	"github.com/alekssro/banking/dto"
	"github.com/alekssro/banking/errs"
)

// AccountService defines a AccountService interface
type AccountService interface {
	CreateAccount(string) (*dto.AccountResponse, *errs.AppError)
}

// DefaultAccountService struct defines the default Account service
// which depends on domain.AccountRepository
type DefaultAccountService struct {
	repo domain.AccountRepository
}

// GetAllAccounts method implements the AccountService interface for
// DefaultAccountService struct
func (s DefaultAccountService) CreateAccount(string) (*dto.AccountResponse, *errs.AppError) {

	response := dto.AccountResponse{}

	return &response, nil
}

// NewAccountService func adds a new default account service
func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
