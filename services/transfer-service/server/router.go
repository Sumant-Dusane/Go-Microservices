package server

import (
	"github.com/gin-gonic/gin"
	"swamisamartha.com/transfer-service/handlers"
)

func NewTransferRouter() *gin.Engine {
	router := gin.New()

	router.GET("/", handlers.GetIndex)
	router.GET("/health", handlers.GetHealth)
	router.GET("/transfer/", handlers.GetTransfers)
	router.GET("/transfer/balance", handlers.GetBalance)

	router.POST("transfer/new", handlers.AddNewTransfer)

	return router
}
