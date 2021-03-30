package db

import (
	"github.com/alekssro/banking/banking/domain/entity"
	"github.com/alekssro/banking/banking/shared/errs"
	"github.com/alekssro/banking/banking/shared/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// CustomerRepositoryDB implements a CustomerRepository
// connector
type CustomerRepositoryDB struct {
	client *sqlx.DB
}

// FindAll method implements a CustomRepository interface for
// CustomerRepositoryDB struct. Returns all the customers
// in CustomerRepositoryDB
func (d CustomerRepositoryDB) FindAll() ([]entity.Customer, *errs.AppError) {

	findAllQuery := "select customer_id, name, date_of_birth, city, zipcode, status from customers"

	var customers []entity.Customer
	// Query + Marshall into customers
	err := d.client.Select(&customers, findAllQuery)
	if err != nil {
		logger.Error("Error while querying customers table in DB. " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

// FindByStatus method implements finding all the customers with a given status
// for the CustomerRepositoryDB
func (d CustomerRepositoryDB) FindByStatus(status string) ([]entity.Customer, *errs.AppError) {

	findByStatusQuery := "select customer_id, name, date_of_birth, city, zipcode, status from customers where status = ?"

	// check valid status in service
	if status != "1" && status != "0" {
		return nil, errs.NewBadRequestError("malformed query, status=" + status)
	}

	var customers []entity.Customer
	// Query by condition + Marshall into customers
	err := d.client.Select(&customers, findByStatusQuery, status)
	if err != nil {
		logger.Error("Error while querying customers table by status in DB. " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

// ByID method implements CustomRepository interface for CustomerRepositoryDB
// struct. Returns a Customer given an ID.
func (d CustomerRepositoryDB) ByID(id string) (*entity.Customer, *errs.AppError) {

	customerQuery := "select customer_id, name, date_of_birth, city, zipcode, status from customers where customer_id = ?"

	var c entity.Customer
	// Query by id + Marshall into c
	err := d.client.Get(&c, customerQuery, id)
	if err != nil {
		logger.Error("Error while querying customer by id in DB. " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &c, nil
}

// NewCustomerRepositoryDB func implements adding a new
// CustomerRepositoryDB client
func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{client: dbClient}
}
