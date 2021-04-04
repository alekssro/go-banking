package application

import (
	"github.com/alekssro/banking/banking/domain/repository"
	"github.com/alekssro/banking/lib/dto"
	"github.com/alekssro/banking/lib/errs"
)

// TransactionService defines a TransactionService interface
type TransactionService interface {
	NewTransaction(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

// DefaultTransactionService struct defines the default Transaction service
// which depends on domain.TransactionRepository
type DefaultTransactionService struct {
	repo repository.TransactionRepository
}

// NewTransactionService func adds a new default Transaction service
func NewTransactionService(repo repository.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repo}
}
