CREATE TABLE IF NOT EXISTS wallets
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(200)   NOT NULL
        CONSTRAINT check_wallet_name_too_short CHECK ( length(name) > 5 ),
    balance    NUMERIC(18, 2) NOT NULL DEFAULT 0.0
        CONSTRAINT check_balance_not_negative CHECK (balance >= 0.0),
    created_at timestamptz    NOT NULL DEFAULT now(),
    updated_at timestamptz    NOT NULL DEFAULT now()
);


CREATE UNIQUE INDEX UNQ_WALLETS_NAME ON wallets (name);

DROP TYPE IF EXISTS UNIT;
CREATE TYPE UNIT AS ENUM ('deposit','transfer','wallet_oper_journal');


CREATE TABLE IF NOT EXISTS wallet_deposits
(
    id         SERIAL PRIMARY KEY,
    wallet     INT            NOT NULL REFERENCES wallets (id),
    amount     NUMERIC(18, 2) NOT NULL,
    created_at timestamptz    NOT NULL DEFAULT now(),
    CHECK ( amount > 0.0 )
);

CREATE TABLE IF NOT EXISTS wallet_transfer
(
    id          SERIAL PRIMARY KEY,
    wallet_from INT            NOT NULL REFERENCES wallets (id),
    wallet_to   INT            NOT NULL REFERENCES wallets (id),
    amount      NUMERIC(18, 2) NOT NULL,
    created_at  timestamptz    NOT NULL DEFAULT now(),
    CHECK ( amount > 0.0 )
);

CREATE TABLE IF NOT EXISTS wallet_oper_journal
(
    id         SERIAL PRIMARY KEY,
    wallet     INT            NOT NULL REFERENCES wallets (id),
    oper_sign  INT            NOT NULL DEFAULT 0, -- 0=+ 1=-
    amount     NUMERIC(18, 2) NOT NULL,
    created_at timestamptz    NOT NULL DEFAULT now(),
    unit       UNIT           NOT NULL,
    CHECK (oper_sign IN (0, 1))
);

CREATE TABLE IF NOT EXISTS wallet_oper_journal_links
(
    id     SERIAL PRIMARY KEY,
    in_id  INT NOT NULL REFERENCES wallet_oper_journal (id),
    out_id INT NOT NULL REFERENCES wallet_oper_journal (id)
);

CREATE UNIQUE INDEX UNQ_WALLETS_OPER_JOURNAL_LINKS ON wallet_oper_journal_links (in_id, out_id);

CREATE TABLE IF NOT EXISTS wallet_unit_links
(
    id       SERIAL PRIMARY KEY,
    in_unit  UNIT NOT NULL,
    out_unit UNIT NOT NULL,
    in_id    INT  NOT NULL,
    out_id   INT  NOT NULL
);

CREATE UNIQUE INDEX UNQ_WALLETS_UNIT_LINKS ON wallet_unit_links (in_unit, in_id, out_unit, out_id);
