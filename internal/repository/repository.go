package repository

import (
	"context"
	"mibshard/configs"
)

type WalletKeeper interface {
	CreateNote(ctx context.Context, key string, value int) error
	SetNote(ctx context.Context, key string, value int) error
	GetNote(ctx context.Context, key string) (string, error)
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
