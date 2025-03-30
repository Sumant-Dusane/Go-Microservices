package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"swamisamartha.com/user-service/db"
	"swamisamartha.com/user-service/server"
)

func main() {
	db.InitDB()
	srv := server.StartHttpServer()

	// Channel to listen for OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Wait for termination signal
	<-stop
	log.Println("Shutting down server...")

	// Gracefully shutdown server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Server shutdown failed: %v", err)
	}

	db.CloseDB()
	log.Println("✅ Server gracefully stopped")
}
