package service

import (
	"context"
	"mibshard/internal/repository"
)

type WalletKeeper interface {
	PrepareTransaction(ctx context.Context, TxID string, walletID int, amount int) error
	CommitChanges(ctx context.Context, TxID string, walletID int, valamountue int) error
	CreateWallet(ctx context.Context, walletID int) error
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
