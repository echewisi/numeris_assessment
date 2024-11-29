package routes

import (
	"net/http"

	"github.com/echewisi/numeris_assessment/internal/handlers"
	"github.com/gorilla/mux"
)

// AuthRouter sets up routes for authentication
func AuthRouter(authHandler *handlers.AuthHandler) *mux.Router {
	router := mux.NewRouter()

	// Authentication routes
	router.HandleFunc("/auth/register", authHandler.RegisterHandler).Methods(http.MethodPost)
	router.HandleFunc("/auth/login", authHandler.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/auth/logout", authHandler.LogoutHandler).Methods(http.MethodPost)

	return router
}
