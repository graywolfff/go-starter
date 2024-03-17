package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// define routers
	http.HandleFunc("/greet", greet)

	// starting server
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func greet (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!!")

}
