package domain

import (
	"strconv"

	"github.com/alekssro/banking/errs"
	"github.com/alekssro/banking/logger"
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
func (d TransactionRepositoryDB) Withdrawal(t Transaction) (*Transaction, *errs.AppError) {

	var a Account

	// 1. Check if available amount in account
	accountAmountQuery := "SELECT amount FROM accounts WHERE account_id = ?"
	err := d.client.Get(&a.Amount, accountAmountQuery, t.AccountID)
	if err != nil {
		logger.Error("Error while getting account amount" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	if !a.CanWithdraw(t.Amount) {
		return nil, errs.NewUnexpectedError("Insufficient funds to make transaction")
	}

	// 2. Insert transaction into transactions
	id, inserr := insertTransaction(d, t)
	if inserr != nil {
		return nil, inserr
	}
	t.TransactionID = strconv.FormatInt(id, 10)

	// 3. Update account amount
	updateAccountQuery := "UPDATE accounts SET amount = ? WHERE account_id = ?"
	newBalance := a.Amount - t.Amount
	_, err = d.client.Exec(updateAccountQuery, newBalance, t.AccountID)
	if err != nil {
		logger.Error("Error while updating account amount: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	t.AccountBalance = newBalance

	return &t, nil
}

// Transaction method implements a TransactionRepository interface for
// TransactionRepositoryDB struct. Returns all the customers
// in TransactionRepositoryDB
func (d TransactionRepositoryDB) Deposit(t Transaction) (*Transaction, *errs.AppError) {

	var a Account

	// 1. Get amount in account
	accountAmountQuery := "SELECT amount FROM accounts WHERE account_id = ?"
	err := d.client.Get(&a.Amount, accountAmountQuery, t.AccountID)
	if err != nil {
		logger.Error("Error while getting account amount" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// 2. Insert transaction into transactions
	id, inserr := insertTransaction(d, t)
	if inserr != nil {
		return nil, inserr
	}
	t.TransactionID = strconv.FormatInt(id, 10)

	// 2. Update account amount
	updateAccountQuery := "UPDATE accounts SET amount = ? WHERE account_id = ?"
	newBalance := a.Amount + t.Amount
	_, err = d.client.Exec(updateAccountQuery, newBalance, t.AccountID)
	if err != nil {
		logger.Error("Error while updating account amount: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	t.AccountBalance = newBalance

	return &t, nil
}

func insertTransaction(d TransactionRepositoryDB, t Transaction) (int64, *errs.AppError) {
	insertTransactionQuery := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES (?, ?, ?, ?)"
	res, err := d.client.Exec(insertTransactionQuery, t.AccountID, t.Amount, t.TransactionType, t.TransactionDate)
	if err != nil {
		logger.Error("Error while inserting account entry: " + err.Error())
		return -1, errs.NewUnexpectedError("unexpected database error")
	}
	id, err := res.LastInsertId()
	if err != nil {
		logger.Error("Error while retrieving last inserted id: " + err.Error())
		return -1, errs.NewUnexpectedError("unexpected database error")
	}
	return id, nil
}

// NewTransactionRepositoryDB func implements adding a new
// TransactionRepositoryDB client
func NewTransactionRepositoryDB(dbClient *sqlx.DB) TransactionRepositoryDB {
	return TransactionRepositoryDB{client: dbClient}
}
