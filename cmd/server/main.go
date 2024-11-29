package main

import (
	"log"
	"net/http"

	"github.com/echewisi/numeris_assessment/internal/config"
	"github.com/echewisi/numeris_assessment/internal/handlers"
	"github.com/echewisi/numeris_assessment/internal/middleware"
	"github.com/echewisi/numeris_assessment/internal/repositories"
	"github.com/echewisi/numeris_assessment/internal/services"
	"github.com/echewisi/numeris_assessment/internal/utils"
	"github.com/echewisi/numeris_assessment/routes"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Initialize DB connection
	db, err := utils.ConnectDatabase(cfg.Database)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Repositories
	invoiceRepo := repositories.NewInvoiceRepository(db)
	customerRepo := repositories.NewCustomerRepository(db)
	paymentRepo := repositories.NewPaymentRepository(db)

	// Services
	invoiceService := services.NewInvoiceService(invoiceRepo)
	customerService := services.NewCustomerService(customerRepo)
	paymentService := services.NewPaymentService(paymentRepo)

	// Handlers
	invoiceHandler := handlers.NewInvoiceHandler(invoiceService)
	customerHandler := handlers.NewCustomerHandler(customerService)
	paymentHandler := handlers.NewPaymentHandler(paymentService)
	authHandler := handlers.NewAuthHandler()

	// Routers
	apiRouter := routes.APIRouter(invoiceHandler, customerHandler, paymentHandler)
	authRouter := routes.AuthRouter(authHandler)

	// Combine routers
	mainRouter := http.NewServeMux()
	mainRouter.Handle("/api/", middleware.AuthMiddleware(apiRouter))
	mainRouter.Handle("/auth/", authRouter)

	// Start server
	serverAddr := ":" + cfg.Server.Port
	log.Printf("Starting server on %s...", serverAddr)
	if err := http.ListenAndServe(serverAddr, mainRouter); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
