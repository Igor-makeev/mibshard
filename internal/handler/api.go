package handler

import (
	"mibshard/internal/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PrepareTransaction(c *gin.Context) {

	var input entities.PreparerRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.Service.WalletKeeper.PrepareTransaction(c.Request.Context(), input.TxID, input.WalletID, input.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusAccepted)

}

func (h *Handler) CommitChanges(c *gin.Context) {

}

func (h *Handler) Createwallet(c *gin.Context) {

	var input entities.CreateWalletRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.WalletKeeper.CreateWallet(c.Request.Context(), input.WalletID, input.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusAccepted)

}
