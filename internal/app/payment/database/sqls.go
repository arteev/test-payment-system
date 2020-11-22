package database

// SQL
const (
	sqlNewWallet               = "insert into wallets(name,balance) values($1,0.0) returning *"
	sqlGetWallet               = "select * from wallets where id = $1"
	sqlDeposit                 = `insert into wallet_deposits(wallet,amount) values($1,$2) returning *`
	sqlInsertWalletOperJournal = `insert into wallet_oper_journal(wallet,oper_sign,amount,unit, hash) 
values($1,$2,$3,$4,$5) returning id`
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
	sqlOperationsWallet            = `select j.id,
       j.wallet wallet_id,
       w.name   wallet_name,
       j.unit,
       j.amount,
       case j.oper_sign
           when 0 then 'deposit'
           when 1 then 'withdraw'
           end  oper_type,
       j.created_at,
       wt.id    wallet_to_id,
       wt.name  wallet_to_name
from wallet_oper_journal j
         left join wallet_oper_journal_links l on j.id = l.in_id
         left join wallet_oper_journal jt on jt.id = l.out_id
         left join wallets wt on wt.id = jt.wallet,
     wallets w
where w.id = j.wallet
  and ($2::INT is null or j.oper_sign = $2)
  and j.wallet = $1
  and ($3::timestamptz is null or j.created_at >= $3)
  and ($4::timestamptz is null or j.created_at <= $4)
order by j.created_at`
	sqlExistsDuplicateOperJournal = `select 1 count
from wallet_oper_journal j
where j.hash = $1
and j.created_at + interval '%s' > now()
limit 1
`
)
