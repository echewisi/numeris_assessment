package models

import "time"

type Activity struct {
	ID        int64     `json:"id" db:"id"`               // Unique ID for the activity
	UserID    int64     `json:"user_id" db:"user_id"`     // Associated user ID
	Action    string    `json:"action" db:"action"`       // Action performed (e.g., login, update invoice)
	Timestamp time.Time `json:"timestamp" db:"timestamp"` // Timestamp of the action
	Metadata  string    `json:"metadata" db:"metadata"`   // Additional info (e.g., IP address, changes made)
}
