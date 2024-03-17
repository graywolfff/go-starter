package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

// create customer struct with json tags
type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
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
	customers := []Customer{}
	customers = append(customers, Customer{
		Name:    "Felix",
		City:    "Ha noi",
		Zipcode: "1000",
	})
	customers = append(customers, Customer{
		Name:    "Phongdh",
		City:    "Hai duong",
		Zipcode: "1110",
	})
	// check the request header for sure what data repesent need to response
	// in this case xml or json
	
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
