package server

import (
	"log"
	"net/http"
)

func StartHttpServer() *http.Server {
	router := NewTransferRouter()

	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal("Error starting transfer-service: ", err)
	}

	return srv
}
