package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alekssro/banking/domain"
	"github.com/alekssro/banking/logger"
	"github.com/alekssro/banking/service"
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
	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDB(dbClient)
	transactionRepositoryDB := domain.NewTransactionRepositoryDB(dbClient)

	ch := CustomerHandler{service.NewCustomerService(customerRepositoryDB)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDB)}
	th := TransactionHandler{service.NewTransactionService(transactionRepositoryDB)}

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
