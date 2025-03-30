package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"swamisamartha.com/transfer-service/db"
)

type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func GetTransfers(c *gin.Context) {
	t, err := db.Get()

	if err != nil {
		err := fmt.Sprintf("Something went wrong: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"success": true, "transfers": t})
}

func GetBalance(c *gin.Context) {
	username, err := getUser()

	t, errr := db.Get()

	if errr != nil {
		err := fmt.Sprintf("Something went wrong: %v", errr)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err != nil {
		err := fmt.Sprintf("Something went wrong: %s", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"user": username, "balance": "1000000000000000$", "history": t})
}

func AddNewTransfer(c *gin.Context) {

	var transfer db.TransferDto

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&transfer); err != nil {
		errMsg := fmt.Sprintf("Unable to parse request body object: %s", err)
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	id, err := db.Add(&transfer)

	if err != nil {
		err := fmt.Sprintf("Something went wrong: %s", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "insertedId": id.InsertedID})
}
