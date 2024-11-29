package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/echewisi/numeris_assessment/internal/models"
	"github.com/echewisi/numeris_assessment/internal/services"
	"github.com/gorilla/mux"
)

type InvoiceHandler struct {
	Service *services.InvoiceService
}

// NewInvoiceHandler creates a new InvoiceHandler
func NewInvoiceHandler(service *services.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{Service: service}
}

// CreateInvoiceHandler handles creating a new invoice
func (h *InvoiceHandler) CreateInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	var invoice models.Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.Service.CreateInvoice(r.Context(), &invoice)
	if err != nil {
		http.Error(w, "Failed to create invoice: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

// GetInvoiceHandler handles fetching an invoice by ID
func (h *InvoiceHandler) GetInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"] // Using mux to get the path parameter
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid invoice ID", http.StatusBadRequest)
		return
	}

	invoice, err := h.Service.GetInvoiceByID(r.Context(), int64(id))
	if err != nil {
		http.Error(w, "Invoice not found: "+err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(invoice)
}

// UpdateInvoiceHandler handles updating an invoice
func (h *InvoiceHandler) UpdateInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid invoice ID", http.StatusBadRequest)
		return
	}

	var invoice models.Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Set the ID from the path to ensure the correct invoice is updated
	invoice.ID = int64(id)

	updatedInvoice, err := h.Service.UpdateInvoice(r.Context(), &invoice)
	if err != nil {
		http.Error(w, "Failed to update invoice: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedInvoice)
}

// DeleteInvoiceHandler handles deleting an invoice
func (h *InvoiceHandler) DeleteInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid invoice ID", http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteInvoice(r.Context(), int64(id)); err != nil {
		http.Error(w, "Failed to delete invoice: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // HTTP 204: No Content
}
