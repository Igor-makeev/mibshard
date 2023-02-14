package entities

type PreparerRequest struct {
	TxID     string `json:"transaction_id" binding:"required"`
	WalletID int    `json:"wallet_id" binding:"required"`
	Amount   int    `json:"amount" binding:"required"`
}
type CreateWalletRequest struct {
	Id int `json:"id" binding:"required"`
}

type ReadyToCommitRequest struct {
	TxID     string `json:"transaction_id" binding:"required"`
	WalletID int    `json:"wallet_id" binding:"required"`
}
