package model

import (
	"database/sql"
	"time"
)

type Operation struct {
	ID            uint           `db:"id"`
	WalletID      uint           `db:"wallet_id"`
	WalletName    string         `db:"wallet_name"`
	Unit          Unit           `db:"unit"`
	CreatedAt     time.Time      `db:"created_at"`
	Amount        float64        `db:"amount"`
	WalletToID    *uint          `db:"wallet_to_id"`
	WalletToName  sql.NullString `db:"wallet_to_name"`
	OperationType string         `db:"oper_type"`
}
