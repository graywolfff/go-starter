package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graywolff/go-banking/domain"
	"github.com/graywolff/go-banking/service"
)

func Start() {
	customer_repo := domain.NewCustomerRepositoryStub()
	customer_service := service.NewCustomerService(&customer_repo)

	ch := CustomerHandlers{
		&customer_service,
	}
	// create a mux server

	mux := http.NewServeMux()
	// define routers
	mux.HandleFunc("GET /customers", ch.getAllCustomers)
	// we can read params on url by using PathValue method of http.Request
	mux.HandleFunc("GET /customers/{id}", func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		fmt.Fprint(w, idString)
	})
	// starting server
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", mux))
}
