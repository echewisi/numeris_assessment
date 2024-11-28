package models

import "time"

type Payment struct {
	ID         int64     `json:"id" db:"id"`               // Unique ID for the payment
	InvoiceID  int64     `json:"invoice_id" db:"invoice_id"` // Associated invoice ID
	Amount     float64   `json:"amount" db:"amount"`       // Payment amount
	Method     string    `json:"method" db:"method"`       // Payment method (e.g., card, cash)
	CreatedAt  time.Time `json:"created_at" db:"created_at"` // Payment creation time
	ConfirmedAt *time.Time `json:"confirmed_at,omitempty" db:"confirmed_at"` // Confirmation time (if confirmed)
}
