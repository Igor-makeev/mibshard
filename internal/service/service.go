package service

import (
	"context"
	"mibshard/internal/repository"
)

type WalletKeeper interface {
	ChangeWalletBalance(ctx context.Context, key int, value int) error
	CreateWallet(ctx context.Context, key int, value int) error
}

type Service struct {
	WalletKeeper
}

func NewService(repo *repository.Repository) *Service {
	return &Service{WalletKeeper: NewWalletKeeperService(repo)}
}
