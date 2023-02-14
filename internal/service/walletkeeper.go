package service

import (
	"context"
	"fmt"
	"mibshard/internal/repository"
)

type WalletKeeperService struct {
	repo *repository.Repository
	tm   *TransactionManager
}

func NewWalletKeeperService(repo *repository.Repository) *WalletKeeperService {
	return &WalletKeeperService{
		repo: repo,
		tm:   NewTransactionManager(repo),
	}
}
func (wks *WalletKeeperService) CreateWallet(ctx context.Context, key int) error {
	if err := wks.repo.CreateWallet(ctx, key); err != nil {
		return err
	}
	return nil
}
func (wks *WalletKeeperService) PrepareTransaction(ctx context.Context, TxID string, walletID int, amount int) error {

	oldBalance, lock_status, err := wks.repo.GetWalletBalance(ctx, walletID)
	if err != nil {
		return err
	}
	if lock_status {

		return &WalletLockedError{walletID}
	}
	newBalance := oldBalance + amount

	if newBalance < 0 {
		err = &InvalidBalanceValueError{walletID}
		return err
	}

	tx := &Transaction{
		txID:        TxID,
		walletID:    walletID,
		amount:      amount,
		prepareFlag: true,
	}
	wks.tm.txLog.AddRecord(ctx, tx.txID, tx.walletID, tx.amount)
	wks.tm.stream <- tx

	return nil
}

func (wks *WalletKeeperService) CommitChanges(ctx context.Context, TxID string, walletID int, amount int) error {

	if err := wks.repo.ChangeWalletBalance(ctx, walletID, amount); err != nil {
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
