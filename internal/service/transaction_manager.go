package service

import (
	"mibshard/internal/repository"
)

type Transaction struct {
	txID string

	amount int

	walletID int

	prepareFlag bool
	commitFlag  bool
}

type TransactionManager struct {
	txLog  repository.TxLogger
	stream chan *Transaction
}

func NewTransactionManager(repo *repository.Repository) *TransactionManager {
	tm := &TransactionManager{
		txLog:  repo,
		stream: make(chan *Transaction, 50),
	}

	return tm
}

// func (tm *TransactionManager) Run(ctx context.Context) {
// 	go tm.listen(ctx)
// }

func (tm *TransactionManager) Close() {
	close(tm.stream)
}

// func (tm *TransactionManager) listen(ctx context.Context) {

// 	for {

// 		select {
// 		case TX := <-tm.stream:

// 		case <-ctx.Done():

// 			return
// 		}
// 	}
// }
