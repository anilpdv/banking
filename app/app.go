package app

import (
	"log"
	"net/http"

	"github.com/anilpdv/banking/domain"
	"github.com/anilpdv/banking/service"
	"github.com/gorilla/mux"
)

// Start of the application
func Start() {
	mux := mux.NewRouter()
	// wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	mux.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}

}
