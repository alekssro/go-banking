package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alekssro/banking/banking/application"
	"github.com/alekssro/banking/banking/domain/repository"
	"github.com/alekssro/banking/banking/infrastructure/persistence/db"
	"github.com/alekssro/banking/lib/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	if os.Getenv("BANKING_HOST") == "" {
		logger.Warn("BANKING_HOST not defined, using default")
		os.Setenv("BANKING_HOST", "0.0.0.0")
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
	// define routes
	router.
		HandleFunc("/customers", ch.queryCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).
		Methods(http.MethodGet).
		Name("GetCustomer")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.createAccount).
		Methods(http.MethodPost).
		Name("NewAccount")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", th.newTransaction).
		Methods(http.MethodPost).
		Name("NewTransaction")

	// authentication middleware layer
	am := AuthMiddleware{repository.NewAuthRepository()}
	router.Use(am.authorizationHandler())

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
