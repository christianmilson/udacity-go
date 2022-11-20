package main

import (
	"net/http"

	"github.com/christianmilson/udacity-go/controllers"
	"github.com/christianmilson/udacity-go/db"
	"github.com/christianmilson/udacity-go/models"
	"github.com/gorilla/mux"
)

func contentTypeApplicationJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func main() {
	db := db.NewDb()
	customerModel := &models.Customer{}

	// Seed the database with 3 records
	db.Seed(3, customerModel)

	controller := controllers.NewCustomerController(db)

	// Register the handlers for the CustomerController
	r := mux.NewRouter()
	r.HandleFunc("/customers", controller.Index).Methods("GET")
	r.HandleFunc("/customers/{id}", controller.Show).Methods("GET")
	r.HandleFunc("/customers/{id}", controller.Delete).Methods("DELETE")
	r.HandleFunc("/customers", controller.Store).Methods("POST")
	r.HandleFunc("/customers/{id}", controller.Update).Methods("PATCH")

	// Automatically add the response json header to all requests
	r.Use(contentTypeApplicationJsonMiddleware)

	// Start the web server on port 8080
	http.ListenAndServe(":8080", r)
}
