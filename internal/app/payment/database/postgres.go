package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	migrations "test-payment-system/internal/app/payment/database/migrate"
	"test-payment-system/internal/app/payment/database/model"
	"test-payment-system/internal/pkg/config"
	"test-payment-system/pkg/database"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres" // comment
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	_ "github.com/jackc/pgx/v4/stdlib" // register pg driver
	"go.uber.org/zap"
)

const duplicateDeltaTime = "2 minutes" // Postgresql interval

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

func (db *DB) insertWalletOperJournal(ctx context.Context, tx *Transaction,
	journal model.WalletOperJournal) (uint, error) {
	var id uint
	err := tx.GetContext(ctx, &id, sqlInsertWalletOperJournal,
		journal.WalletID,
		journal.OperSign,
		journal.Amount,
		journal.Unit,
		journal.Hash)
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
		var err error
		if returnErr != nil {
			err = tx.Rollback()
		} else {
			err = tx.Commit()
		}
		if err != nil {
			db.Logger.Errorw("tx error", zap.Error(err))
		}
	}()

	err = tx.GetContext(ctx, deposit, sqlDeposit, walletID, amount)
	if err != nil {
		return nil, processPgError(err, "deposit wallet")
	}

	journalItem := model.WalletOperJournal{
		WalletID: walletID,
		OperSign: model.OperationSignIncome,
		Amount:   amount,
		Unit:     model.UnitNameDeposit,
	}
	journalItem.Hash = journalItem.GetHashWalletOperation()
	if err = db.checkDuplicateWalletOperJournal(ctx, tx, journalItem); err != nil {
		return nil, processPgError(err, "")
	}
	journalID, err := db.insertWalletOperJournal(ctx, tx, journalItem)
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
		var err error
		if returnErr != nil {
			err = tx.Rollback()
		} else {
			err = tx.Commit()
		}
		if err != nil {
			db.Logger.Errorw("tx error", zap.Error(err))
		}
	}()

	err = tx.GetContext(ctx, transfer, sqlTransfer, walletFrom, walletTo, amount)
	if err != nil {
		return nil, processPgError(err, "deposit wallet")
	}

	journalItem := model.WalletOperJournal{
		WalletID: walletFrom,
		OperSign: model.OperationSignTransfer,
		Amount:   amount,
		Unit:     model.UnitNameTransfer,
	}
	journalItem.Hash = journalItem.GetHashWalletOperation(walletTo)

	err = db.checkDuplicateWalletOperJournal(ctx, tx, journalItem)
	if err != nil {
		return nil, processPgError(err, "")
	}

	journalFromID, err := db.insertWalletOperJournal(ctx, tx, journalItem)
	if err != nil {
		return nil, processPgError(err, "")
	}

	journalItem = model.WalletOperJournal{
		WalletID: walletTo,
		OperSign: model.OperationSignIncome,
		Amount:   amount,
		Unit:     model.UnitNameTransfer,
	}
	journalItem.Hash = journalItem.GetHashWalletOperation()
	journalToID, err := db.insertWalletOperJournal(ctx, tx, journalItem)
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

func (db *DB) checkDuplicateWalletOperJournal(ctx context.Context, tx *Transaction,
	journal model.WalletOperJournal) error {
	if journal.Hash == "" {
		return nil
	}

	var exists int
	sqlCommand := fmt.Sprintf(sqlExistsDuplicateOperJournal, duplicateDeltaTime) // problem sqlx parameters with interval?
	err := tx.GetContext(ctx, &exists, sqlCommand, journal.Hash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return processPgError(err, "")
	}
	if exists > 0 {
		return fmt.Errorf("%w: period %s", ErrDuplicateWalletOperation, duplicateDeltaTime)
	}

	return nil
}

func (db *DB) OperationWallet(ctx context.Context, walletID uint, operSign *model.OperationSign, timeFrom,
	timeTo time.Time) ([]*model.Operation, error) {
	connection := db.PgDatabase.Connection
	var timeFromArg, timeToArg *time.Time
	if !timeFrom.IsZero() {
		timeFromArg = &timeFrom
	}
	if !timeTo.IsZero() {
		timeToArg = &timeTo
	}
	rows, err := connection.QueryxContext(ctx, sqlOperationsWallet, walletID, operSign, timeFromArg, timeToArg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, processPgError(err, "wallet_oper_journal")
	}
	defer rows.Close()

	operations := make([]*model.Operation, 0)
	for rows.Next() {
		operation := &model.Operation{}
		err := rows.StructScan(operation)
		if err != nil {
			return nil, err
		}
		operations = append(operations, operation)
	}
	return operations, nil
}
