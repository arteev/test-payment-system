package database

import (
	"context"
	"github.com/jmoiron/sqlx"
	"test-payment-system/internal/app/payment/database/model"
)

// Transaction alias of type sqlx.Tx
type Transaction = sqlx.Tx

// Database wallet repository
type Database interface {
	NewWallet(ctx context.Context, name string) (*model.Wallet, error)
	GetWallet(ctx context.Context, id uint) (*model.Wallet, error)
	Deposit(ctx context.Context, walletID uint, amount float64) (*model.WalletDeposit, error)
}
