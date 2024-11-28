package models

import "time"

type Invoice struct {
	ID          int64     `json:"id" db:"id"`               // Unique ID for the invoice
	CustomerID  int64     `json:"customer_id" db:"customer_id"` // ID of the customer
	Amount      float64   `json:"amount" db:"amount"`       // Invoice amount
	Status      string    `json:"status" db:"status"`       // Status (e.g., paid, unpaid)
	CreatedAt   time.Time `json:"created_at" db:"created_at"` // Invoice creation time
	DueDate     time.Time `json:"due_date" db:"due_date"`   // Due date for payment
	PaidAt      *time.Time `json:"paid_at,omitempty" db:"paid_at"` // Payment date (if paid)
	Description string    `json:"description" db:"description"` // Optional description
}
