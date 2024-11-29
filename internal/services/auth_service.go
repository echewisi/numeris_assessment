package services

import (
	"context"
	"errors"
	"time"

	"github.com/echewisi/numeris_assessment/internal/models"
	"github.com/echewisi/numeris_assessment/internal/repositories"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo          *repositories.UserRepository
	JWTSecretKey  string
	TokenDuration time.Duration
}

// NewAuthService creates a new AuthService
func NewAuthService(repo *repositories.UserRepository, jwtSecretKey string, tokenDuration time.Duration) *AuthService {
	return &AuthService{Repo: repo, JWTSecretKey: jwtSecretKey, TokenDuration: tokenDuration}
}

// RegisterUser registers a new user
func (s *AuthService) RegisterUser(ctx context.Context, user *models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	return s.Repo.CreateUser(ctx, user)
}

// Login authenticates the user and generates a JWT token
func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.Repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(s.TokenDuration).Unix(),
	})

	return token.SignedString([]byte(s.JWTSecretKey))
}

// Logout invalidates the JWT token (implementation depends on your setup)
func (s *AuthService) Logout(ctx context.Context, token string) error {
	// You can implement token revocation using a blacklist or cache mechanism.
	return nil
}
