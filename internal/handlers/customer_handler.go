package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/echewisi/numeris_assessment/internal/models"
	"github.com/echewisi/numeris_assessment/internal/services"
)

type CustomerHandler struct {
	Service *services.CustomerService
}

// NewCustomerHandler creates a new CustomerHandler
func NewCustomerHandler(service *services.CustomerService) *CustomerHandler {
	return &CustomerHandler{Service: service}
}

// CreateCustomerHandler handles creating a new customer
func (h *CustomerHandler) CreateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.Service.CreateCustomer(r.Context(), &customer)
	if err != nil {
		http.Error(w, "Failed to create customer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

// GetCustomerHandler handles fetching a customer by ID
func (h *CustomerHandler) GetCustomerHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	customer, err := h.Service.GetCustomer(r.Context(), int64(id))
	if err != nil {
		http.Error(w, "Customer not found: "+err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(customer)
}

// UpdateCustomerHandler handles updating a customer
func (h *CustomerHandler) UpdateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.Service.UpdateCustomer(r.Context(), &customer)
	if err != nil {
		http.Error(w, "Failed to update customer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// DeleteCustomerHandler handles deleting a customer by ID
func (h *CustomerHandler) DeleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteCustomer(r.Context(), int64(id))
	if err != nil {
		http.Error(w, "Failed to delete customer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Customer deleted successfully"))
}
