package server

import (
	"github.com/gin-gonic/gin"
	"swamisamartha.com/transfer-service/handlers"
)

func NewTransferRouter() *gin.Engine {
	router := gin.New()

	router.GET("/", handlers.GetIndex)
	router.GET("/health", handlers.GetHealth)
	router.GET("/transfer/balance", handlers.GetBalance)

	return router
}
