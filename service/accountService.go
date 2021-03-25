package service

import (
	"time"

	"github.com/alekssro/banking/domain"
	"github.com/alekssro/banking/dto"
	"github.com/alekssro/banking/errs"
)

// AccountService defines a AccountService interface
type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

// DefaultAccountService struct defines the default Account service
// which depends on domain.AccountRepository
type DefaultAccountService struct {
	repo domain.AccountRepository
}

// GetAllAccounts method implements the AccountService interface for
// DefaultAccountService struct
func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToNewAccountDTO()

	return &response, nil
}

// NewAccountService func adds a new default account service
func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
