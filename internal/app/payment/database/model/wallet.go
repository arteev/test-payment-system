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

// WalletDeposit wallet deposit model
type WalletDeposit struct {
	ID        uint      `db:"id"`
	WalletID  uint      `db:"wallet"`
	Amount    float64   `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
}
