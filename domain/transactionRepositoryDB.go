package domain

import (
	"github.com/alekssro/banking/errs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// TransactionRepositoryDB implements a TransactionRepository connector
type TransactionRepositoryDB struct {
	client *sqlx.DB
}

// Transaction method implements a TransactionRepository interface for
// TransactionRepositoryDB struct. Returns all the customers
// in TransactionRepositoryDB
func (d TransactionRepositoryDB) Transaction(a Transaction) (*Transaction, *errs.AppError) {
	// TODO: Implement
	return &a, nil
}

// NewTransactionRepositoryDB func implements adding a new
// TransactionRepositoryDB client
func NewTransactionRepositoryDB(dbClient *sqlx.DB) TransactionRepositoryDB {
	return TransactionRepositoryDB{client: dbClient}
}
