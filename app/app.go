package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alekssro/banking/domain"
	"github.com/alekssro/banking/logger"
	"github.com/alekssro/banking/service"
	"github.com/gorilla/mux"
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
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// define routes
	router.HandleFunc("/customers", ch.queryCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	host := os.Getenv("BANKING_HOST")
	port := os.Getenv("BANKING_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router))

}
