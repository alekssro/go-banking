package application

import (
	"time"

	"github.com/alekssro/banking/banking/domain/entity"
	"github.com/alekssro/banking/lib/dto"
	"github.com/alekssro/banking/lib/errs"
)

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
		response := toNewTransactionResponseDTO(*newTransaction)
		return &response, nil
	}
	// deposit
	newTransaction, err := s.repo.Deposit(t)
	if err != nil {
		return nil, err
	}
	response := toNewTransactionResponseDTO(*newTransaction)
	return &response, nil

}

func toNewTransactionResponseDTO(t entity.Transaction) dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		TransactionId:  t.TransactionID,
		AccountBalance: t.AccountBalance,
	}
}
