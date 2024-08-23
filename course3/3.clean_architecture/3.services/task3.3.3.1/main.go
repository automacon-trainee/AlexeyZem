package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"projectService/internal/controller"
)

func main() {
	server := http.Server{
		Addr:         ":8080",
		Handler:      controller.NewRouter(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("starting server")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	sig := <-sigs
	log.Println("Got signal:", sig, "start shutting down")
	timeForShutDown := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeForShutDown)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Error while shutting down:", err)
	}
	time.Sleep(timeForShutDown)
	log.Println("Server exiting")
}
