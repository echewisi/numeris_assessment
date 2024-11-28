package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTConfig struct {
	SecretKey string
	Issuer    string
	ExpiresIn time.Duration
}

// Claims defines the structure of JWT claims
type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

// GenerateToken generates a JWT token for a given user
func GenerateToken(config JWTConfig, userID, role string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.ExpiresIn).Unix(),
			Issuer:    config.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SecretKey))
}

// ValidateToken parses and validates a JWT token
func ValidateToken(config JWTConfig, tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
