package app

import (
	"log"
	"net/http"

	"github.com/alekssro/banking/domain"
	"github.com/alekssro/banking/service"
	"github.com/gorilla/mux"
)

// Start the web app
func Start() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryMock())}
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
