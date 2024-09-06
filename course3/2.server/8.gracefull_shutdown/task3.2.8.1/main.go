// Реализуй API-сервер на языке программирования Golang с применением graceful shutdown. Сервер должен
// обрабатывать HTTP-запросы и корректно завершать свою работу при получении сигнала остановки.

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

	"auth/internal"
)

func main() {
	server := http.Server{
		Addr:         ":8080",
		Handler:      internal.NewRouter(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	sig := <-sigs
	log.Println("Got signal:", sig, "shutting down")
	timeForShutDown := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeForShutDown)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Error while shutting down:", err)
	}
	time.Sleep(timeForShutDown)
	log.Println("Server exiting")
}
