package handler

import (
	"mibshard/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
	Router  *gin.Engine
}

func NewHandler(service *service.Service) *Handler {
	handler := &Handler{
		Router:  gin.New(),
		Service: service,
	}

	api := handler.Router.Group("/api")
	{

		api.POST("/changevalue", handler.ChangeWalletBalance)
		api.POST("/createwallet", handler.Createwallet)

	}

	return handler
}
