package server

import (
	"log"
	"net/http"
)

func StartHttpServer() *http.Server {
	router := NewUserServiceRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error in starting server: ", err)
	}

	return srv
}
