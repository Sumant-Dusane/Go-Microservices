package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	c.JSON(http.StatusAccepted, " || SHREE SWAMI SAMARTHA || ")
}

func GetHealth(c *gin.Context) {
	c.JSON(http.StatusAccepted, "Transfer service is healthy!.")
}
