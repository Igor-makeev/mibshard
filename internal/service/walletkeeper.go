package service

import (
	"context"
	"fmt"
	"mibshard/internal/repository"
)

type WalletKeeperService struct {
	repo *repository.Repository
}

func NewWalletKeeperService(repo *repository.Repository) *WalletKeeperService {
	return &WalletKeeperService{repo: repo}
}
func (wks *WalletKeeperService) CreateWallet(ctx context.Context, key int, value int) error {
	if err := wks.repo.CreateWallet(ctx, key, value); err != nil {
		return err
	}
	return nil
}
func (wks *WalletKeeperService) ChangeWalletBalance(ctx context.Context, key int, value int) error {

	oldBalance, err := wks.repo.GetWalletBalance(ctx, key)
	if err != nil {
		return err
	}

	newBalance := oldBalance + value
	if newBalance < 0 {
		err = &InvalidBalanceValueError{newBalance}
		return err
	}

	if err = wks.repo.ChangeWalletBalance(ctx, key, newBalance); err != nil {
		return err
	}
	return nil
}

type InvalidBalanceValueError struct {
	value int
}

func (ibve *InvalidBalanceValueError) Error() string {
	return fmt.Sprintf("error: new ballance %v has invalid value", ibve.value)
}
