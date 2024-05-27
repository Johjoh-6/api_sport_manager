package models

import (
	"time"
)

// Bill represents a bill
type Bill struct {
	ID        string    `json:"id,omitempty"`
	TeamID    int       `json:"team_id"`
	Amount    float64   `json:"amount"`
	Currency  string    `json:"currency"`
	BillDate  time.Time `json:"bill_date"`
	IDPayment string    `json:"id_payment"`
}
