package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// create customer struct with json tags
type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

func main() {

	// define routers
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)
	// starting server
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!!")

}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{
			Name:    "Felix",
			City:    "Ha Noi",
			Zipcode: "10000",
		},
	}
	json.NewEncoder(w).Encode(customers)
}
