package database

import "github.com/jmoiron/sqlx"

// Transaction type for sqlx.Tx
type Transaction = sqlx.Tx

type Database interface {

}