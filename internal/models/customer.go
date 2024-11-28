package models

import "time"

type Customer struct {
	ID        int64     `json:"id" db:"id"`               // Unique ID for the customer
	Name      string    `json:"name" db:"name"`           // Customer's name
	Email     string    `json:"email" db:"email"`         // Customer's email
	Phone     string    `json:"phone" db:"phone"`         // Phone number
	Address   string    `json:"address" db:"address"`     // Physical address
	CreatedAt time.Time `json:"created_at" db:"created_at"` // Account creation time
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"` // Last update time
}
