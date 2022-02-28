package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"fullname"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func Start() {

	router := mux.NewRouter()

	// GET /greet
	router.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!!")
	}).Methods(http.MethodGet)

	// GET /customers
	router.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		customers := []Customer{
			{Name: "Ashich", City: "New Delhi", Zipcode: "110075"},
			{Name: "Rob", City: "New Delhi", Zipcode: "110075"},
		}

		w.Header().Add("Content-Type", "application/json")

		json.NewEncoder(w).Encode(customers)
	}).Methods(http.MethodGet)

	// POST /customers
	router.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Post method received")
	}).Methods(http.MethodPost)

	// GET /customers/{customer_id:[0-9]+}
	router.HandleFunc("/customers/{customer_id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		fmt.Fprint(w, vars["customer_id"])
	}).Methods(http.MethodGet)

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
