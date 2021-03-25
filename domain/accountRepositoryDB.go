package domain

import (
	"strconv"

	"github.com/alekssro/banking/errs"
	"github.com/alekssro/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// AccountRepositoryDB implements a AccountRepository connector
type AccountRepositoryDB struct {
	client *sqlx.DB
}

// FindAll method implements a AccountRepository interface for
// AccountRepositoryDB struct. Returns all the customers
// in AccountRepositoryDB
func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {

	sqlInsertQuery := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"

	res, err := d.client.Exec(sqlInsertQuery, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while inserting account entry: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	id, err := res.LastInsertId()
	if err != nil {
		logger.Error("Error while retrieving last inserted id: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

// NewAccountRepositoryDB func implements adding a new
// AccountRepositoryDB client
func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{client: dbClient}
}
