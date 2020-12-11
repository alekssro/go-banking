package app

import (
	"log"
	"net/http"
)

// Start the web app
func Start() {
	mux := http.NewServeMux()

	// define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))

}
