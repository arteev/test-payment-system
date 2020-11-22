ALTER TABLE wallet_oper_journal
    ADD COLUMN IF NOT EXISTS hash varchar(64);