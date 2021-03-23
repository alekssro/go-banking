package domain

import (
	"database/sql"
	"time"

	"github.com/alekssro/banking/errs"
	"github.com/alekssro/banking/logger"
	_ "github.com/go-sql-driver/mysql"
)

// CustomerRepositoryDB implements a CustomerRepository
// connector
type CustomerRepositoryDB struct {
	client *sql.DB
}

// FindAll method implements a CustomRepository interface for
// CustomerRepositoryDB struct. Returns all the customers
// in CustomerRepositoryDB
func (d CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppError) {

	findAllQuery := "select customer_id, name, date_of_birth, city, zipcode, status from customers"

	rows, err := d.client.Query(findAllQuery)
	if err != nil {
		logger.Error("Error while logging into DB." + err.Error())
		return nil, errs.NewUnexpectedError("Error while logging into DB.")
	}

	var customers []Customer
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.DateofBirth, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("Customer not found")
			} else {
				logger.Error("Error while scanning customer in DB")
				return nil, errs.NewUnexpectedError("Unexpected database error")
			}
		}
		customers = append(customers, c)
	}

	return customers, nil
}

// FindByStatus method implements finding all the customers with a given status
// for the CustomerRepositoryDB
func (d CustomerRepositoryDB) FindByStatus(status string) ([]Customer, *errs.AppError) {

	if status != "1" && status != "0" {
		return nil, errs.NewBadRequestError("malformed query, status=" + status)
	}

	findByStatusQuery := "select customer_id, name, date_of_birth, city, zipcode, status from customers where status = ?"

	// Query by condition
	rows, err := d.client.Query(findByStatusQuery, status)
	if err != nil {
		logger.Error("Error while logging into DB." + err.Error())
		return nil, errs.NewUnexpectedError("Error while logging into DB.")
	}

	var customers []Customer
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.DateofBirth, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("Customer not found")
			} else {
				logger.Error("Error while scanning customer in DB")
				return nil, errs.NewUnexpectedError("Unexpected database error")
			}
		}
		customers = append(customers, c)
	}

	return customers, nil
}

// ByID method implements CustomRepository interface for CustomerRepositoryDB
// struct. Returns a Customer given an ID.
func (d CustomerRepositoryDB) ByID(id string) (*Customer, *errs.AppError) {
	cutomerQuery := "select customer_id, name, date_of_birth, city, zipcode, status from customers where customer_id = ?"

	// Looking for one row
	row := d.client.QueryRow(cutomerQuery, id)

	// Error handling
	var c Customer
	err := row.Scan(&c.ID, &c.Name, &c.DateofBirth, &c.City, &c.Zipcode, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer in DB")
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &c, nil
}

// NewCustomerRepositoryDB func implements adding a new
// CustomerRepositoryDB client
func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root:codecamp@tcp(mysql-dev:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client: client}
}
