package database

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"test-payment-system/internal/pkg/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // comment
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jackc/pgx/v4/stdlib" // register pg driver
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const defaultDriverPGX = "pgx"

type ResourcesGetter interface {
	GetResources() ([]string, bindata.AssetFunc)
}

// DB Base Database
type PgDatabase struct {
	Connection *sqlx.DB
	Logger     *zap.SugaredLogger
	config     config.DBConfig
	resource   ResourcesGetter
}

func New(cfg *config.DBConfig, log *zap.SugaredLogger,
	resource ResourcesGetter) (*PgDatabase, error) {
	dsn := cfg.URI
	connection, err := open(log, defaultDriverPGX, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to db open: %w", err)
	}
	db := &PgDatabase{
		Logger:     log,
		Connection: connection,
		config:     *cfg,
		resource:   resource,
	}
	if err := db.Healthcheck(context.TODO()); err != nil {
		return nil, fmt.Errorf("failed to health check: %w", err)
	}

	if err := db.MigrateUp(); err != nil {
		return nil, fmt.Errorf("failed to migrate: %w", err)
	}
	log.Debug("DB connected.")
	return db, nil
}

func open(log *zap.SugaredLogger, driver, dsn string) (*sqlx.DB, error) {
	connStr := dsn
	if driver == defaultDriverPGX {
		config, err := pgx.ParseConfig(dsn)
		if err != nil {
			return nil, err
		}
		config.Logger = newLoggerWrapPgx(log)
		config.LogLevel = getLevelFromZap(log)
		connStr = stdlib.RegisterConnConfig(config)
	}

	connection, err := sqlx.Open(driver, connStr)
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func (db *PgDatabase) Healthcheck(ctx context.Context) error {
	return db.Connection.DB.PingContext(ctx)
}

func (db *PgDatabase) MigrateUp() error {
	// Prepare resources
	dsn, err := getURIMigrate(db.config)
	if err != nil {
		return err
	}
	names, assets := db.resource.GetResources()
	resources := bindata.Resource(names, assets)
	driver, err := bindata.WithInstance(resources)
	if err != nil {
		return err
	}

	migrateInstance, err := migrate.NewWithSourceInstance("go-bindata", driver, dsn)
	if err != nil {
		return err
	}
	defer migrateInstance.Close()

	err = migrateInstance.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	migrated := err == nil
	version, dirty, err := migrateInstance.Version()
	if err != nil {
		db.Logger.Error(err)
	} else {
		message := "Migrated to version DB: %d, dirty: %v"
		if !migrated {
			message = "Current version DB: %d, dirty: %v"
		}
		db.Logger.Infof(message, version, dirty)
	}
	return nil
}

func (db *PgDatabase) MigrateDown() error {
	// Prepare resources
	dsn, err := getURIMigrate(db.config)
	if err != nil {
		return err
	}
	names, assets := db.resource.GetResources()
	resources := bindata.Resource(names, assets)
	driver, err := bindata.WithInstance(resources)
	if err != nil {
		return err
	}

	migrateInstance, err := migrate.NewWithSourceInstance("go-bindata", driver, dsn)
	if err != nil {
		return err
	}
	defer migrateInstance.Close()

	version, dirty, err := migrateInstance.Version()
	if err != nil {
		return err
	}
	if dirty {
		return errors.New("database dirty")
	}
	db.Logger.Debugw("migrate down", "current_version", version)
	err = migrateInstance.Down()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

// getURIMigrate corrects the uri for the migration register and returns it
func getURIMigrate(cfg config.DBConfig) (string, error) {
	result := cfg.URIMigrate
	if result == "" {
		result = cfg.URI
	}
	if cfg.ForceTableMigrations != "" {
		const keyMigrateTable = "x-migrations-table"
		uri, err := url.Parse(result)
		if err != nil {
			return "", err
		}
		values := uri.Query()
		if values.Get(keyMigrateTable) != "" {
			return result, nil
		}
		values.Set(keyMigrateTable, cfg.ForceTableMigrations)
		uri.RawQuery = values.Encode()
		result = uri.String()
	}
	return result, nil
}

func (db *PgDatabase) Close() error {
	return db.Connection.Close()
}
