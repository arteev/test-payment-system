package model

import "time"

// Wallet model
type Wallet struct {
	ID        uint      `db:"id"`
	Name      string    `db:"name"`
	Balance   float64   `db:"balance"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
