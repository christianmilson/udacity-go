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
	// Seed the database with 3 records
	db.Seed(3, &models.Customer{})

	customerController := controllers.NewCustomerController(db)
	homeController := controllers.NewHomeController(db)

	r := mux.NewRouter()
	r.HandleFunc("/", homeController.Index).Methods("GET")

	// Register the handlers for the CustomerController
	customerRouter := r.PathPrefix("/customers").Subrouter()
	customerRouter.HandleFunc("", customerController.Index).Methods("GET")
	customerRouter.HandleFunc("/{id}", customerController.Show).Methods("GET")
	customerRouter.HandleFunc("/{id}", customerController.Delete).Methods("DELETE")
	customerRouter.HandleFunc("", customerController.Store).Methods("POST")
	customerRouter.HandleFunc("/{id}", customerController.Update).Methods("PATCH")

	// Automatically add the response json header to all requests
	customerRouter.Use(contentTypeApplicationJsonMiddleware)

	// Start the web server on port 8080
	http.ListenAndServe(":8080", r)
}
