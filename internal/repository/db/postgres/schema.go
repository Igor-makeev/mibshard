package repository

const walletKeeperSchema = `
create table if not exists wallet_keeper (
    ID integer NOT NULL UNIQUE,
    Balance integer,
  Lock_Status boolean
);`

const transactionLogSchema = `
create table if not exists transaction_log (
    Transaction_id varchar (40),
    Wallet_id integer,
    Amount integer,
    Status varchar (10)
	);`
