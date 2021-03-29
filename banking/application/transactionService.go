package application

import (
	"time"

	"github.com/alekssro/banking/banking/domain/entity"
	"github.com/alekssro/banking/banking/shared/dto"
	"github.com/alekssro/banking/banking/shared/errs"
)

// TransactionService defines a TransactionService interface
type TransactionService interface {
	NewTransaction(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

// DefaultTransactionService struct defines the default Transaction service
// which depends on domain.TransactionRepository
type DefaultTransactionService struct {
	repo entity.NewTransactionRepository
}

// GetAllTransactions method implements the TransactionService interface for
// DefaultTransactionService struct
func (s DefaultTransactionService) NewTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	t := entity.Transaction{
		AccountID:       req.AccountID,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	if t.TransactionType == "withdrawal" {
		newTransaction, err := s.repo.Withdrawal(t)
		if err != nil {
			return nil, err
		}
		response := newTransaction.ToTransactionDTO()
		return &response, nil
	}
	// deposit
	newTransaction, err := s.repo.Deposit(t)
	if err != nil {
		return nil, err
	}
	response := newTransaction.ToTransactionDTO()
	return &response, nil

}

// NewTransactionService func adds a new default Transaction service
func NewTransactionService(repo entity.NewTransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repo}
}
