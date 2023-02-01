package repository

import (
	"context"
	"mibshard/configs"
)

type WalletKeeper interface {
	CreateWallet(ctx context.Context, key int, value int) error
	ChangeWalletBalance(ctx context.Context, key int, value int) error
	GetWalletBalance(ctx context.Context, key int) (int, error)
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
