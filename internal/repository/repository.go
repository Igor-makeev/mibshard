package repository

import (
	"context"
	"mibshard/configs"
)

type PersistanceStorage interface {
	WalletKeeper
	TxLogger
}

type WalletKeeper interface {
	CreateWallet(ctx context.Context, walletID int) error
	ChangeWalletBalance(ctx context.Context, walletID int, amount int) error
	GetWalletBalance(ctx context.Context, walletID int) (int, bool, error)
	Close(ctx context.Context) error
}

type TxLogger interface {
	AddRecord(ctx context.Context, TxId string, walletID int, amount int) error
	ChangeStatus(ctx context.Context, TxId, status string) error
}
type Repository struct {
	PersistanceStorage
	cfg *configs.Config
}

func NewRepository(ps PersistanceStorage, cfg *configs.Config) *Repository {
	return &Repository{
		PersistanceStorage: ps,
		cfg:                cfg,
	}
}

func (rep *Repository) Close(ctx context.Context) error {
	return rep.PersistanceStorage.Close(ctx)
}
