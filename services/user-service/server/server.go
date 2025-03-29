package server

import (
	"log"
	"net/http"
)

func InitUserService() {
	startHttpServer()
}

func startHttpServer() *http.Server {
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
