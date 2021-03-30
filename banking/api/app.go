package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alekssro/banking/banking/application"
	"github.com/alekssro/banking/banking/infrastructure/persistence/db"
	"github.com/alekssro/banking/banking/shared/logger"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	if os.Getenv("BANKING_HOST") == "" {
		logger.Warn("BANKING_HOST not defined, using default")
		os.Setenv("BANKING_HOST", "localhost")
	}
	if os.Getenv("BANKING_PORT") == "" {
		logger.Warn("BANKING_PORT not defined, using default")
		os.Setenv("BANKING_PORT", "3000")
	}
}

// Start the web app
func Start() {

	// Sanity Check
	sanityCheck()

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// wiring
	dbClient := getDBclient()

	customerRepositoryDB := db.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := db.NewAccountRepositoryDB(dbClient)
	transactionRepositoryDB := db.NewTransactionRepositoryDB(dbClient)

	ch := CustomerHandler{application.NewCustomerService(customerRepositoryDB)}
	ah := AccountHandler{application.NewAccountService(accountRepositoryDB)}
	th := TransactionHandler{application.NewTransactionService(transactionRepositoryDB)}

	// define routes
	router.HandleFunc("/customers", ch.queryCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.createAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", th.newTransaction).Methods(http.MethodPost)

	// starting server
	host := os.Getenv("BANKING_HOST")
	port := os.Getenv("BANKING_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router))

}

func getDBclient() *sqlx.DB {
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
	return client
}
