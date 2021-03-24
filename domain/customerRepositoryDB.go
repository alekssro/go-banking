package domain

import (
	"fmt"
	"os"
	"time"

	"github.com/alekssro/banking/errs"
	"github.com/alekssro/banking/logger"
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
func (d CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppError) {

	findAllQuery := "select customer_id, name, date_of_birth, city, zipcode, status from customers"

	var customers []Customer
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
func (d CustomerRepositoryDB) FindByStatus(status string) ([]Customer, *errs.AppError) {

	findByStatusQuery := "select customer_id, name, date_of_birth, city, zipcode, status from customers where status = ?"

	if status != "1" && status != "0" {
		return nil, errs.NewBadRequestError("malformed query, status=" + status)
	}

	var customers []Customer
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
func (d CustomerRepositoryDB) ByID(id string) (*Customer, *errs.AppError) {

	customerQuery := "select customer_id, name, date_of_birth, city, zipcode, status from customers where customer_id = ?"

	var c Customer
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
func NewCustomerRepositoryDB() CustomerRepositoryDB {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)

	client, err := sqlx.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client: client}
}
