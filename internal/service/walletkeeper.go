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
func (wks *WalletKeeperService) CreateWallet(ctx context.Context, key int) error {
	if err := wks.repo.CreateWallet(ctx, key); err != nil {
		return err
	}
	return nil
}
func (wks *WalletKeeperService) ChangeWalletBalance(ctx context.Context, key int, value int) error {

	oldBalance, lock_status, err := wks.repo.GetWalletBalance(ctx, key)
	if err != nil {
		return err
	}
	if lock_status {

		return &WalletLockedError{key}
	}
	newBalance := oldBalance + value

	if newBalance < 0 {
		err = &InvalidBalanceValueError{key}
		return err
	}

	if err = wks.repo.ChangeWalletBalance(ctx, key, newBalance); err != nil {
		return err
	}
	return nil
}

type WalletLockedError struct {
	value int
}

func (ibve *WalletLockedError) Error() string {
	return fmt.Sprintf("error: wallet:%v is locked", ibve.value)
}

type InvalidBalanceValueError struct {
	value int
}

func (ibve *InvalidBalanceValueError) Error() string {
	return fmt.Sprintf("error: there is not enough money in the wallet:%v", ibve.value)
}

func (wks *WalletKeeperService) Close(ctx context.Context) error {
	return wks.repo.Close(ctx)
}
