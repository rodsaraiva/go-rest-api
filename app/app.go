package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodsaraiva/go-rest-api/domain"
	"github.com/rodsaraiva/go-rest-api/service"
)

func Start() {

	router := mux.NewRouter()

	// wiring
	// handlers := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	handlers := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// GET /customers
	router.HandleFunc("/customers", handlers.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", handlers.getCustomer).Methods(http.MethodGet)

	// GET /health
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "OK") }).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
