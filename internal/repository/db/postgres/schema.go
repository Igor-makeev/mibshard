package repository

const walletKeeperSchema = `
create table if not exists wallet_keeper (
  user_id integer not null,  
  wallet_id integer NOT NULL UNIQUE,
    Balance integer
  
);`

const transactionLogSchema = `
create table if not exists transaction_log (
    Transaction_id varchar (40),
    Wallet_id integer,
    Amount integer,
    Status varchar (10)
	);`
