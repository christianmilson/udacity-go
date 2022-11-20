package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/christianmilson/udacity-go/db"
	"github.com/christianmilson/udacity-go/models"
	"github.com/gorilla/mux"
)

type CustomerController struct {
	db *db.Database
}

func getId(r *http.Request) int {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	return id
}

func NewCustomerController(db *db.Database) *CustomerController {
	return &CustomerController{
		db: db,
	}
}

func (controller *CustomerController) Index(w http.ResponseWriter, r *http.Request) {
	// Get records
	records := controller.db.FetchAll()

	response := []models.Customer{}

	// Remove the keys
	for _, record := range records {
		response = append(response, record)
	}

	w.WriteHeader(http.StatusOK)

	js := map[string][]models.Customer{
		"data": response,
	}

	json.NewEncoder(w).Encode(js)
}

func (controller *CustomerController) Store(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate that the data is acceptable
	validation := validate(customer)

	// If there are validation errors => display to user and set 422 header
	if len(validation) > 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		js := map[string]map[string]string{
			"errors": validation,
		}

		json.NewEncoder(w).Encode(js)

		return
	}

	controller.db.Insert(customer)

	js := map[string]string{
		"message": "CustomerResource created",
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(js)
}

func (controller *CustomerController) Show(w http.ResponseWriter, r *http.Request) {
	record := controller.db.FindById(getId(r))
	if record.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{
			"message": "CustomerResource not found",
		}

		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusOK)
		response := map[string]models.Customer{
			"data": record,
		}

		json.NewEncoder(w).Encode(response)
	}
}

func (controller *CustomerController) Update(w http.ResponseWriter, r *http.Request) {
	// Check that the model actually exists
	if controller.db.FindById(getId(r)).Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{
			"message": "CustomerResource not found",
		}

		json.NewEncoder(w).Encode(response)

		return
	}

	var customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate that the data is acceptable
	validation := validate(customer)

	// If there are validation errors => display to user and set 422 header
	if len(validation) > 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		js := map[string]map[string]string{
			"errors": validation,
		}

		json.NewEncoder(w).Encode(js)

		return
	}

	// Update the existing record into the db
	customer.Id = getId(r)
	controller.db.Update(customer)

	w.WriteHeader(http.StatusOK)
	js := map[string]string{
		"message": "CustomerResource updated",
	}

	json.NewEncoder(w).Encode(js)
}

func (controller *CustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	// Check that the record actually exists
	id := getId(r)
	if controller.db.FindById(id).Id == 0 {
		js := map[string]string{
			"errors": "CustomerResource not found",
		}

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(js)

		return
	}

	// Delete the record
	controller.db.Delete(getId(r))

	response := map[string]string{
		"message": "CustomerResource deleted",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func validate(customer models.Customer) map[string]string {
	emailRegex := "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	errors := map[string]string{}

	// Basic Validation
	if customer.Email == "" {
		errors["email"] = "Email cannot be empty"
	}

	if !regexp.MustCompile(emailRegex).MatchString(customer.Email) {
		errors["email"] = "Email is not a valid email"
	}

	if customer.Name == "" {
		errors["name"] = "Name cannot be empty"
	}

	if customer.Role == "" {
		errors["role"] = "Role cannot be empty"
	}

	if customer.Phone == "" {
		errors["phone"] = "Phone cannot be empty"
	}

	return errors
}
