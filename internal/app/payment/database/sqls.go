package database

// SQL
const (
	sqlNewWallet = "insert into wallets(name,balance) values($1,0.0) returning *"
	sqlGetWallet = "select * from wallets where id = $1"
)
