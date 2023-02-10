package repository

const postgresSchema = `
create table if not exists wallet_keeper (
    ID integer NOT NULL UNIQUE,
    Balance integer,
  Lock_Status boolean
);`
