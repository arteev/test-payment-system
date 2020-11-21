package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
)

var (
	ErrWrongBalance         = errors.New("balance must not be negative")
	ErrWalletNameIsTooShort = errors.New("wallet name is too short")
	ErrWalletNameMustUnique = errors.New("wallet name must be unique")
)

func processPgError(err error, entityName string) error {
	if err == nil {
		return nil
	}
	switch vErr := err.(type) {
	default:
		if err == sql.ErrNoRows {
			if entityName == "" {
				return ErrNotFound
			}
			return fmt.Errorf("%w: %s", ErrNotFound, entityName)
		}
		return err
	case *pgconn.PgError:
		if vErr.ConstraintName != "" {
			return processPgConstraintError(vErr)
		}
	}
	return err
}

func processPgConstraintError(pgError *pgconn.PgError) error {
	switch pgError.ConstraintName {
	default:
		return pgError
	case "check_balance_not_negative":
		return ErrWrongBalance
	case "check_wallet_name_too_short":
		return ErrWalletNameIsTooShort
	case "unq_wallets_name":
		return ErrWalletNameMustUnique
	}

}
