package service

import (
	"context"
	"mibshard/internal/repository"
)

type WalletKeeper interface {
	ChangeWalletBalance(ctx context.Context, key int, value int) error
	CreateWallet(ctx context.Context, key int) error
	Close(ctx context.Context) error
}

type Service struct {
	WalletKeeper
}

func NewService(repo *repository.Repository) *Service {
	return &Service{WalletKeeper: NewWalletKeeperService(repo)}
}

func (srv *Service) Close(ctx context.Context) error {
	return srv.WalletKeeper.Close(ctx)
}
