package entities

type ChangeRequest struct {
	Id     string `json:"id" binding:"required"`
	Amount int    `json:"amount" binding:"required"`
}
type CreateWalletRequest struct {
	Id string `json:"id" binding:"required"`
}
