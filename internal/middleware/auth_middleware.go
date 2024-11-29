package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/echewisi/numeris_assessment/internal/utils"
	"github.com/echewisi/numeris_assessment/pkg/config"
)

type ContextKey string

const UserContextKey ContextKey = "user"

// AuthMiddleware validates the JWT token in the Authorization header
func AuthMiddleware(cfg *config.Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized: Missing or invalid token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Convert config.JWTConfig to utils.JWTConfig
		jwtConfig := utils.JWTConfig{
			SecretKey: cfg.JWT.SecretKey,
			Issuer:    cfg.JWT.Issuer,
			ExpiresIn: cfg.JWT.ExpiresIn,
		}

		// Validate the token
		claims, err := utils.ValidateToken(jwtConfig, tokenString)
		if err != nil {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Add user data from claims to the context
		ctx := context.WithValue(r.Context(), UserContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

