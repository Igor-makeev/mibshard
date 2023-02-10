package repository

import (
	"context"
	"mibshard/configs"
)

type WalletKeeper interface {
	CreateWallet(ctx context.Context, key int) error
	ChangeWalletBalance(ctx context.Context, key int, value int) error
	GetWalletBalance(ctx context.Context, key int) (int, bool, error)
	Close(ctx context.Context) error
}
type Repository struct {
	WalletKeeper
	cfg *configs.Config
}

func NewRepository(wk WalletKeeper, cfg *configs.Config) *Repository {
	return &Repository{
		WalletKeeper: wk,
		cfg:          cfg,
	}
}

func (rep *Repository) Close(ctx context.Context) error {
	return rep.WalletKeeper.Close(ctx)
}
