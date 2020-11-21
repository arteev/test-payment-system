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

// OperationSign journal operation sign
type OperationSign int

// Known operation sign
const (
	OperationSignIncome   OperationSign = 0
	OperationSignTransfer OperationSign = 1
)

type WalletOperJournal struct {
	ID       uint          `db:"uint"`
	WalletID uint          `db:"wallet"`
	OperSign OperationSign `db:"oper_sign"`
	Amount   float64       `db:"amount"`
	Unit     Unit          `db:"unit"`
}

// WalletTransfer wallet transfer money model
type WalletTransfer struct {
	ID         uint      `db:"id"`
	WalletFrom uint      `db:"wallet_from"`
	WalletTo   uint      `db:"wallet_to"`
	Amount     float64   `db:"amount"`
	CreatedAt  time.Time `db:"created_at"`
}
