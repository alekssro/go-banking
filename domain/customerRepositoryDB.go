package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// CustomerRepositoryDB implements a CustomerRepository
// connector
type CustomerRepositoryDB struct {
	client *sql.DB
}

// FindAll implements a CustomRepository interface for
// CustomerRepositoryDB struct. Returns all the customers
// in CustomerRepositoryDB
func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {

	findAll_query := "select customer_id, name, date_of_birth, city, zipcode, status from customers"

	rows, err := d.client.Query(findAll_query)
	if err != nil {
		log.Println("Error while logging into DB.", err.Error())
		return nil, err
	}

	var customers []Customer
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.DateofBirth, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			log.Println("Error while scanning customer in DB.", err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}

	return customers, nil
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
