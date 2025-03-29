package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"swamisamartha.com/shared"
)

func GetBalance(c *gin.Context) {
	username, err := getUser()
	if err != nil {
		err := fmt.Sprintf("Something went wrong: %s", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"user": username, "balance": "1000000000000000$"})
}

func getUser() (*shared.User, error) {
	url := "http://user_service:8080/auth/user"

	res, err := http.Get(url)

	if err != nil {
		log.Printf("[Transfer Service > getUser]: %s", err.Error())
		return nil, fmt.Errorf("error getting user from user-service")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Printf("[Transfer Service > getUser]: %s", err.Error())
		return nil, fmt.Errorf("error while reading response")
	}

	var user shared.User

	if err := json.Unmarshal(body, &user); err != nil {
		log.Printf("[Transfer Service > getUser]: %s", err.Error())
		return nil, fmt.Errorf("error while decoding user")
	}

	return &user, nil
}
