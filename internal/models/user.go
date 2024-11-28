package models

import "time"

type User struct {
	ID        int64     `json:"id" db:"id"`             // Unique ID for the user
	Email     string    `json:"email" db:"email"`       // Email address
	Password  string    `json:"-" db:"password"`        // Hashed password
	Role      string    `json:"role" db:"role"`         // Role (e.g., admin, user)
	CreatedAt time.Time `json:"created_at" db:"created_at"` // Account creation time
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"` // Last update time
}
