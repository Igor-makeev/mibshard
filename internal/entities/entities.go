package entities

type ChangeRequest struct {
	Id     int `json:"id" binding:"required"`
	Amount int `json:"amount" binding:"required"`
}
type CreateWalletRequest struct {
	Id int `json:"id" binding:"required"`
}
