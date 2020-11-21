package database

// SQL
const (
	sqlNewWallet               = "insert into wallets(name,balance) values($1,0.0) returning *"
	sqlGetWallet               = "select * from wallets where id = $1"
	sqlDeposit                 = `insert into wallet_deposits(wallet,amount) values($1,$2) returning *`
	sqlInsertWalletOperJournal = `insert into wallet_oper_journal(wallet,oper_sign,amount,unit) 
values($1,$2,$3,$4) returning id`
	sqlInsertUnitLink      = `insert into wallet_unit_links(in_unit,out_unit,in_id,out_id) values($1,$2,$3,$4) returning id`
	sqlUpdateBalanceWallet = `update wallets w
set 
	updated_at = now(),
	balance = (
    select coalesce(sum(case
                            when j.oper_sign = 0 then 1
                            else -1 end * j.amount), 0) :: numeric(18, 2) balance
    from wallet_oper_journal j
    where j.wallet = $1)
where w.id = $1`
	sqlTransfer                    = `insert into wallet_transfer(wallet_from,wallet_to,amount) values($1,$2,$3) returning *`
	sqlInsertWalletOperJournalLink = `insert into  wallet_oper_journal_links(in_id,out_id) values ($1,$2) returning id`
)
