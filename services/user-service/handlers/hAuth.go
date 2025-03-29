package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"swamisamartha.com/shared"
)

func GetUser(c *gin.Context) {
	user := shared.User{
		Id:       "123",
		Name:     "Sumant",
		Email:    "sumant.dusane1966@gmail.com",
		Nickname: "c0mrade",
	}
	c.JSON(http.StatusOK, user)
}
