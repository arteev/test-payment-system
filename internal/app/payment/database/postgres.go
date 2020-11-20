package database

import (
	"context"
	"database/sql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // comment
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	_ "github.com/jackc/pgx/v4/stdlib" // register pg driver
	"go.uber.org/zap"
	migrations "test-payment-system/internal/app/payment/database/migrate"
	"test-payment-system/internal/pkg/config"
	"test-payment-system/pkg/database"
)

const defaultDriverSQLX = "pgx"

// Errors
var (
	ErrNotFound = sql.ErrNoRows
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

// Start returns new transaction for db
func (db *DB) Start(ctx context.Context) (*Transaction, error) {
	return db.PgDatabase.Connection.Beginx()
}

