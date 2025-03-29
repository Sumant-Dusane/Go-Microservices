package server

import (
	"github.com/gin-gonic/gin"

	"swamisamartha.com/user-service/handlers"
)

func NewUserServiceRouter() *gin.Engine {
	router := gin.New()

	router.GET("/", handlers.GetIndex)
	router.GET("/health", handlers.GetHealth)
	router.GET("/auth/user", handlers.GetUser)

	return router
}
