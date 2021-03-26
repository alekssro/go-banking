package service

import (
	"time"

	"github.com/alekssro/banking/domain"
	"github.com/alekssro/banking/dto"
	"github.com/alekssro/banking/errs"
)

// TransactionService defines a TransactionService interface
type TransactionService interface {
	NewTransaction(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

// DefaultTransactionService struct defines the default Transaction service
// which depends on domain.TransactionRepository
type DefaultTransactionService struct {
	repo domain.NewTransactionRepository
}

// GetAllTransactions method implements the TransactionService interface for
// DefaultTransactionService struct
func (s DefaultTransactionService) NewTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	t := domain.Transaction{
		TransactionID:   "",
		AccountID:       req.AccountID,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	newTransaction, err := s.repo.Transaction(t)
	if err != nil {
		return nil, err
	}
	response := newTransaction.ToTransactionDTO()

	return &response, nil
}

// NewTransactionService func adds a new default Transaction service
func NewTransactionService(repo domain.NewTransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repo}
}
