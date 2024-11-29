package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"gorm.io/gorm"


	"github.com/echewisi/numeris_assessment/internal/models"
	"github.com/echewisi/numeris_assessment/internal/services"
)

type PaymentHandler struct {
	Service *services.PaymentService
}

// NewPaymentHandler creates a new PaymentHandler
func NewPaymentHandler(service *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{Service: service}
}

// CreatePaymentHandler handles creating a new payment
func (h *PaymentHandler) CreatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.Service.CreatePayment(r.Context(), &payment)
	if err != nil {
		http.Error(w, "Failed to create payment: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

	// UpdatePaymentHandler handles updating a payment
func (h *PaymentHandler) UpdatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.Service.UpdatePayment(r.Context(), &payment)
	if err != nil {
		http.Error(w, "Failed to update payment: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}


// GetPaymentByIDHandler handles retrieving a payment by its ID
func (h *PaymentHandler) GetPaymentByIDHandler(w http.ResponseWriter, r *http.Request) {
    // Extract the "id" parameter from the URL query string
    idParam := r.URL.Query().Get("id")
    id, err := strconv.ParseInt(idParam, 10, 64)
    if err != nil {
        http.Error(w, "Invalid payment ID: "+err.Error(), http.StatusBadRequest)
        return
    }

    // Call the service layer to fetch the payment
    payment, err := h.Service.GetPayment(r.Context(), id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            http.Error(w, "Payment not found", http.StatusNotFound)
        } else {
            http.Error(w, "Failed to retrieve payment: "+err.Error(), http.StatusInternalServerError)
        }
        return
    }

    // Write the payment details to the response
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(payment)
}


// DeletePaymentHandler handles deleting a payment by ID
func (h *PaymentHandler) DeletePaymentHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeletePayment(r.Context(), int64(id))
	if err != nil {
		http.Error(w, "Failed to delete payment: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Payment deleted successfully"))
}


