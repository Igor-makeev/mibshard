package repository

import (
	"context"
	"errors"
	"fmt"
	"mibshard/configs"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type PostgresWalletKeeper struct {
	DB  *pgx.Conn
	cfg configs.Config
	sync.Mutex
}

func NewPostgresWalletKeeper(cfg *configs.Config, conn *pgx.Conn) *PostgresWalletKeeper {

	ps := &PostgresWalletKeeper{
		DB:  conn,
		cfg: *cfg,
	}
	return ps
}

func NewPostgresClient(cfg *configs.Config) (*pgx.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	conn, err := pgx.Connect(ctx, cfg.PostgresAddres)
	if err != nil {
		logrus.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}
	_, err = conn.Exec(context.Background(), postgresSchema)

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

type WalletNotExistError struct {
	value int
}

func (ibve *WalletNotExistError) Error() string {
	return fmt.Sprintf("error: wallet with this ID-%v does not exist", ibve.value)
}

func (pwk *PostgresWalletKeeper) Close(ctx context.Context) error {
	pwk.DB.Close(ctx)
	return nil
}
