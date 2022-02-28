package app

import (
	"encoding/json"
	"net/http"

	"github.com/rodsaraiva/go-rest-api/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, _ := ch.service.GetAllCustomer()

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(customers)
}
