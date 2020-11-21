package database

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // comment
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	_ "github.com/jackc/pgx/v4/stdlib" // register pg driver
	"go.uber.org/zap"
	migrations "test-payment-system/internal/app/payment/database/migrate"
	"test-payment-system/internal/app/payment/database/model"
	"test-payment-system/internal/pkg/config"
	"test-payment-system/pkg/database"
)

const defaultDriverSQLX = "pgx"

// Errors
var (
	ErrNotFound = errors.New("entity not found")
)

// DB type for db
type DB struct {
	*database.PgDatabase
}

var _ Database = (*DB)(nil)

var _ database.ResourcesGetter = (*DB)(nil)

// New create instance of db
func New(cfg *config.DBConfig, log *zap.SugaredLogger) (*DB, error) {
	db := &DB{}
	pgDB, err := database.New(cfg, log, db)
	if err != nil {
		return nil, err
	}
	db.PgDatabase = pgDB
	return db, nil
}

// GetResources returns bindata resources
func (db *DB) GetResources() ([]string, bindata.AssetFunc) {
	return migrations.AssetNames(), migrations.Asset
}

// Start returns new transaction for db. Default isolation level.
func (db *DB) Start(ctx context.Context) (*Transaction, error) {
	return db.PgDatabase.Connection.BeginTxx(ctx, nil)
}

func (db *DB) NewWallet(ctx context.Context, name string) (*model.Wallet, error) {
	newWallet := &model.Wallet{}
	connection := db.PgDatabase.Connection
	err := connection.GetContext(ctx, newWallet, sqlNewWallet, name)
	if err != nil {
		return nil, processPgError(err, "wallet")
	}
	return newWallet, nil
}

func (db *DB) GetWallet(ctx context.Context, id uint) (*model.Wallet, error) {
	wallet := &model.Wallet{}
	connection := db.PgDatabase.Connection
	err := connection.GetContext(ctx, wallet, sqlGetWallet, id)
	if err != nil {
		return nil, processPgError(err, "wallet")
	}
	return wallet, nil
}

func (db *DB) insertWalletOperJournal(ctx context.Context, tx *Transaction, journal model.WalletOperJournal) (uint, error) {
	var id uint
	err := tx.GetContext(ctx, &id, sqlInsertWalletOperJournal,
		journal.WalletID,
		journal.OperSign,
		journal.Amount,
		journal.Unit)
	if err != nil {
		return 0, processPgError(err, "wallet_oper_journal")
	}
	return id, nil
}

func (db *DB) insertJournalLink(ctx context.Context, tx *Transaction, in, out uint) (uint, error) {
	var id uint
	err := tx.GetContext(ctx, &id, sqlInsertWalletOperJournalLink, in, out)
	if err != nil {
		return 0, processPgError(err, "wallet_oper_journal_links")
	}
	return id, nil
}

func (db *DB) insertUnitLink(ctx context.Context, tx *Transaction, link model.UnitLink) (uint, error) {
	var id uint
	err := tx.GetContext(ctx, &id, sqlInsertUnitLink, link.In, link.Out, link.InID, link.OutID)
	if err != nil {
		return 0, processPgError(err, "wallet_unit_link")
	}
	return id, nil
}

func (db *DB) calculateAndUpdateBalanceWallet(ctx context.Context, tx *Transaction, walletID uint) error {
	result, err := tx.ExecContext(ctx, sqlUpdateBalanceWallet, walletID)
	if err != nil {
		return err
	}
	if affected, err := result.RowsAffected(); err != nil {
		return processPgError(err, "")
	} else if affected == 0 {
		return fmt.Errorf("%w: %s", ErrNotFound, "wallet")
	}
	return nil
}

func (db *DB) Deposit(ctx context.Context, walletID uint, amount float64) (_ *model.WalletDeposit, returnErr error) {
	deposit := &model.WalletDeposit{}
	tx, err := db.Start(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if returnErr != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = tx.GetContext(ctx, deposit, sqlDeposit, walletID, amount)
	if err != nil {
		return nil, processPgError(err, "deposit wallet")
	}

	journalID, err := db.insertWalletOperJournal(ctx, tx, model.WalletOperJournal{
		WalletID: walletID,
		OperSign: model.OperationSignIncome,
		Amount:   amount,
		Unit:     model.UnitNameDeposit,
	})
	if err != nil {
		return nil, processPgError(err, "")
	}

	_, err = db.insertUnitLink(ctx, tx, model.UnitLink{
		In:    model.UnitNameDeposit,
		Out:   model.UnitNameWalletOperJournal,
		InID:  deposit.ID,
		OutID: journalID,
	})
	if err != nil {
		return nil, processPgError(err, "")
	}

	if err = db.calculateAndUpdateBalanceWallet(ctx, tx, walletID); err != nil {
		return nil, processPgError(err, "")
	}

	return deposit, nil
}

func (db *DB) Transfer(ctx context.Context, walletFrom, walletTo uint,
	amount float64) (_ *model.WalletTransfer, returnErr error) {
	transfer := &model.WalletTransfer{}
	tx, err := db.Start(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if returnErr != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = tx.GetContext(ctx, transfer, sqlTransfer, walletFrom, walletTo, amount)
	if err != nil {
		return nil, processPgError(err, "deposit wallet")
	}

	journalFromID, err := db.insertWalletOperJournal(ctx, tx, model.WalletOperJournal{
		WalletID: walletFrom,
		OperSign: model.OperationSignTransfer,
		Amount:   amount,
		Unit:     model.UnitNameTransfer,
	})
	if err != nil {
		return nil, processPgError(err, "")
	}

	journalToID, err := db.insertWalletOperJournal(ctx, tx, model.WalletOperJournal{
		WalletID: walletTo,
		OperSign: model.OperationSignIncome,
		Amount:   amount,
		Unit:     model.UnitNameTransfer,
	})
	if err != nil {
		return nil, processPgError(err, "")
	}

	_, err = db.insertJournalLink(ctx, tx, journalFromID, journalToID)
	if err != nil {
		return nil, processPgError(err, "")
	}

	_, err = db.insertUnitLink(ctx, tx, model.UnitLink{
		In:    model.UnitNameTransfer,
		Out:   model.UnitNameWalletOperJournal,
		InID:  transfer.ID,
		OutID: journalFromID,
	})
	if err != nil {
		return nil, processPgError(err, "")
	}

	_, err = db.insertUnitLink(ctx, tx, model.UnitLink{
		In:    model.UnitNameTransfer,
		Out:   model.UnitNameWalletOperJournal,
		InID:  transfer.ID,
		OutID: journalToID,
	})
	if err != nil {
		return nil, processPgError(err, "")
	}

	if err = db.calculateAndUpdateBalanceWallet(ctx, tx, walletFrom); err != nil {
		return nil, processPgError(err, "")
	}
	if err = db.calculateAndUpdateBalanceWallet(ctx, tx, walletTo); err != nil {
		return nil, processPgError(err, "")
	}
	return transfer, nil
}
