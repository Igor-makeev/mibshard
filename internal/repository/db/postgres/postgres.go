package repository

import (
	"context"
	"errors"
	"fmt"
	"mibshard/configs"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type PostgresWalletKeeper struct {
	DB  *pgxpool.Pool
	cfg configs.Config
}

func NewPostgresWalletKeeper(cfg *configs.Config, conn *pgxpool.Pool) *PostgresWalletKeeper {

	ps := &PostgresWalletKeeper{
		DB:  conn,
		cfg: *cfg,
	}
	return ps
}

func NewPostgresClient(cfg *configs.Config) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	conn, err := pgxpool.New(ctx, cfg.DBAddress)
	if err != nil {
		logrus.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}
	_, err = conn.Exec(context.Background(), walletKeeperSchema)
	logrus.Print(err)
	_, err = conn.Exec(context.Background(), transactionLogSchema)
	logrus.Print(err)
	return conn, err
}

func (pwk *PostgresWalletKeeper) CreateWallet(ctx context.Context, key int) error {
	if _, err := pwk.DB.Exec(ctx, "insert into wallet_keeper(id,balance,lock_status) values($1,$2,$3);", key, 0, false); err != nil {
		return err
	}
	return nil
}
func (pwk *PostgresWalletKeeper) ChangeWalletBalance(ctx context.Context, key int, value int) error {

	if _, err := pwk.DB.Exec(ctx, "update wallet_keeper set balance =$1 where id =$2;", value, key); err != nil {
		return err
	}
	return nil
}
func (pwk *PostgresWalletKeeper) GetWalletBalance(ctx context.Context, key int) (int, bool, error) {
	var result struct {
		balance     int
		lock_status bool
	}
	if err := pwk.DB.QueryRow(ctx, "select balance, lock_status from wallet_keeper where id=$1", key).Scan(&result.balance, &result.lock_status); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = &WalletNotExistError{key}
		}
		return 0, false, err

	}

	return result.balance, result.lock_status, nil

}

func (pwk *PostgresWalletKeeper) AddRecord(ctx context.Context, TxId string, walletID int, amount int) error {
	if _, err := pwk.DB.Exec(ctx, "insert into transaction_log(Transaction_id,Wallet_id,Amount,Status) values($1,$2,$3,$4);", TxId, walletID, amount, "PROCESSING"); err != nil {
		return err
	}
	return nil

}

func (pwk *PostgresWalletKeeper) ChangeStatus(ctx context.Context, TxId, status string) error {

	if _, err := pwk.DB.Exec(ctx, "update transaction_log set Status =$1 where Transaction_id =$2;", TxId, status); err != nil {
		return err
	}
	return nil
}

type WalletNotExistError struct {
	value int
}

func (ibve *WalletNotExistError) Error() string {
	return fmt.Sprintf("error: wallet with this ID-%v does not exist", ibve.value)
}

func (pwk *PostgresWalletKeeper) Close(ctx context.Context) error {
	pwk.DB.Close()
	return nil
}
