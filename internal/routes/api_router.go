package routes

import (
	"net/http"

	"github.com/echewisi/numeris_assessment/internal/handlers"
	"github.com/gorilla/mux"
)

// APIRouter sets up the routes for the API
func APIRouter(invoiceHandler *handlers.InvoiceHandler, customerHandler *handlers.CustomerHandler, paymentHandler *handlers.PaymentHandler) *mux.Router {
	router := mux.NewRouter()

	// Invoice routes
	router.HandleFunc("/invoices", invoiceHandler.CreateInvoiceHandler).Methods(http.MethodPost)
	router.HandleFunc("/invoices/{id:[0-9]+}", invoiceHandler.GetInvoiceHandler).Methods(http.MethodGet)
	router.HandleFunc("/invoices/{id:[0-9]+}", invoiceHandler.UpdateInvoiceHandler).Methods(http.MethodPut)
	router.HandleFunc("/invoices/{id:[0-9]+}", invoiceHandler.DeleteInvoiceHandler).Methods(http.MethodDelete)

	// Customer routes
	router.HandleFunc("/customers", customerHandler.CreateCustomerHandler).Methods(http.MethodPost)
	router.HandleFunc("/customers/{id:[0-9]+}", customerHandler.GetCustomerHandler).Methods(http.MethodGet)

	// Payment routes
	router.HandleFunc("/payments", paymentHandler.CreatePaymentHandler).Methods(http.MethodPost)
	router.HandleFunc("/payments/{id:[0-9]+}", paymentHandler.GetPaymentByIDHandler).Methods(http.MethodGet)

	return router
}
