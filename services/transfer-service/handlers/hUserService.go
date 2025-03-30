package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Microservice communication

func getUser() (*User, error) {
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

	var user User

	if err := json.Unmarshal(body, &user); err != nil {
		log.Printf("[Transfer Service > getUser]: %s", err.Error())
		return nil, fmt.Errorf("error while decoding user")
	}

	return &user, nil
}
