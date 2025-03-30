package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func GetUser(c *gin.Context) {
	user := User{
		Id:       "123",
		Name:     "Sumant",
		Email:    "sumant.dusane1966@gmail.com",
		Nickname: "c0mrade",
	}
	c.JSON(http.StatusOK, user)
}
