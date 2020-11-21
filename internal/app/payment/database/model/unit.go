package model

type Unit string

// Known unit name
const (
	UnitNameDeposit           Unit = "deposit"
	UnitNameTransfer          Unit = "transfer"
	UnitNameWalletOperJournal Unit = "wallet_oper_journal"
)

// UnitLink wallent unit link model
type UnitLink struct {
	ID    uint `db:"id"`
	In    Unit `db:"in_unit"`
	Out   Unit `db:"out_unit"`
	InID  uint `db:"in_id"`
	OutID uint `db:"out_id"`
}
