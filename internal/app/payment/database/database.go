package database

import (
	"context"
	"test-payment-system/internal/app/payment/database/model"
	"time"

	"github.com/jmoiron/sqlx"
)

// Transaction alias of type sqlx.Tx
type Transaction = sqlx.Tx

// Database wallet repository
type Database interface {
	NewWallet(ctx context.Context, name string) (*model.Wallet, error)
	GetWallet(ctx context.Context, id uint) (*model.Wallet, error)
	Deposit(ctx context.Context, walletID uint, amount float64) (*model.WalletDeposit, error)
	Transfer(ctx context.Context, walletFrom, walletTo uint, amount float64) (*model.WalletTransfer, error)
	OperationWallet(ctx context.Context, walletID uint, operSign *model.OperationSign, timeFrom,
		timeTo time.Time) ([]*model.Operation, error)
}

type Migrater interface {
	MigrateDown() error
	MigrateUp() error
}
