package service

import (
	"context"
	"fmt"
	"mibshard/internal/repository"
	"strconv"
)

type WalletKeeperService struct {
	repo *repository.Repository
}

func NewWalletKeeperService(repo *repository.Repository) *WalletKeeperService {
	return &WalletKeeperService{repo: repo}
}
func (wks *WalletKeeperService) CreateWallet(ctx context.Context, key string, value int) error {
	if err := wks.repo.CreateNote(ctx, key, value); err != nil {
		return err
	}
	return nil
}
func (wks *WalletKeeperService) SetNote(ctx context.Context, key string, value int) error {

	oldBalanceString, err := wks.repo.GetNote(ctx, key)
	if err != nil {
		return err
	}

	oldBalance, err := strconv.Atoi(oldBalanceString)
	if err != nil {
		return err
	}

	newBalance := oldBalance + value
	if newBalance < 0 {
		err = &InvalidBalanceValueError{newBalance}
		return err
	}

	if err = wks.repo.SetNote(ctx, key, newBalance); err != nil {
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
