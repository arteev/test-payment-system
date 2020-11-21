package database

// SQL
const (
	sqlNewWallet = "insert into wallets(name,balance) values($1,0.0) returning *"
	sqlGetWallet = "select * from wallets where id = $1"
	sqlDeposit = `insert into wallet_deposits(wallet,amount) values($1,$2) returning *`
)
