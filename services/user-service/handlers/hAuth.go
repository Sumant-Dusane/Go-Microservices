package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"swamisamartha.com/user-service/db"
)

type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func GetUser(c *gin.Context) {
	// user, err := db.Get()
	user, err := db.GetById("551137c2f9e1fac808a5f572")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
