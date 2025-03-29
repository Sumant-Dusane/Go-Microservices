package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {
	c.JSON(http.StatusAccepted, 3000)
}
