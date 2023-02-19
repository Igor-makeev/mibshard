package entities

type PreparerRequest struct {
	TxID     string `json:"transaction_id" binding:"required"`
	WalletID int    `json:"wallet_id" binding:"required"`
	Amount   int    `json:"amount" binding:"required"`
}
type CreateWalletRequest struct {
	WalletID int `json:"wallet_id" binding:"required"`
	UserID   int `json:"user_id" binding:"required"`
}

type txDTO struct {
	TxID     string `json:"transaction_id" binding:"required"`
	WalletID int    `json:"wallet_id" binding:"required"`
}
