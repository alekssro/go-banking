package application

import (
	"time"

	"github.com/alekssro/banking/banking/domain/entity"
	"github.com/alekssro/banking/banking/domain/repository"
	"github.com/alekssro/banking/banking/shared/dto"
	"github.com/alekssro/banking/banking/shared/errs"
)

// GetAllAccounts method implements the AccountService interface for
// DefaultAccountService struct
func (s DefaultAccountService) CreateAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	a := entity.Account{
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
	response := presentNewAccountDTO(*newAccount)

	return &response, nil
}

// NewAccountService func adds a new default account service
func NewAccountService(repo repository.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}

func presentNewAccountDTO(a entity.Account) dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}
