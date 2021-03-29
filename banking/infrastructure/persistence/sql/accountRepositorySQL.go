package persistance

import (
	"strconv"

	"github.com/alekssro/banking/banking/domain/entity"
	"github.com/alekssro/banking/banking/shared/errs"
	"github.com/alekssro/banking/banking/shared/logger"
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
func (d AccountRepositoryDB) Save(a entity.Account) (*entity.Account, *errs.AppError) {

	sqlInsertQuery := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"

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

func (d AccountRepositoryDB) FindByID(accountID string) (*entity.Account, *errs.AppError) {
	// Check if available amount in account
	accountQuery := "SELECT account_id, customer_id, opening_date, account_type, amount, status FROM customers WHERE customer_id = ?"

	var a entity.Account
	// Query by id + Marshall into c
	err := d.client.Get(&a, accountQuery, accountID)
	if err != nil {
		logger.Error("Error while querying customer by id in DB. " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &a, nil
}

// NewAccountRepositoryDB func implements adding a new
// AccountRepositoryDB client
func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{client: dbClient}
}
