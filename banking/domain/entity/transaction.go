package entity

// Transaction struct defines the domain transaction entity
type Transaction struct {
	TransactionID   string `db:"transaction_id"`
	TransactionType string `db:"transaction_type"`
	TransactionDate string `db:"transaction_date"`
	Amount          float64
	AccountID       string `db:"account_id"`
	AccountBalance  float64
}
