package handler

import (
	"mibshard/internal/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ChangeWalletBalance(c *gin.Context) {

	var input entities.ChangeRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.Service.WalletKeeper.ChangeWalletBalance(c.Request.Context(), input.Id, input.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusAccepted)

}

func (h *Handler) Createwallet(c *gin.Context) {

	var input entities.CreateWalletRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.Service.WalletKeeper.CreateWallet(c.Request.Context(), input.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusAccepted)

}
