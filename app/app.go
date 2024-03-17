package app

import (
	"log"
	"net/http"
)

func Start() {
	// create a mux server

	mux := http.NewServeMux()
	// define routers
	mux.HandleFunc("GET /greet", greet)
	mux.HandleFunc("GET /customers", getAllCustomers)

	// starting server
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", mux))
}
