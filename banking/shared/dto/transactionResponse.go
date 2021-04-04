package dto

// Transaction struct defines the Transaction for API response
type NewTransactionResponse struct {
	TransactionId  string  `json:"transaction_id"`
	AccountBalance float64 `json:"balance"`
}
