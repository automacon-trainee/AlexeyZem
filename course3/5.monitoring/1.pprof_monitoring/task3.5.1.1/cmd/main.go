// Подключи профайлер pprof к веб-серверу геосервиса и проанализируй полученные данные.
// Измени  стандартные  роуты  профайлера  так,  чтобы  они  были  доступны  только  авторизованным
// пользователям.
// Ресурс url для получения данных по профилированию: /mycustompath/pprof

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pprof/internal/server"
)

func main() {
	serv, err := server.NewServer()
	if err != nil {
		panic(err)
	}
	go func() {
		log.Println("Server start with addr:", serv.Addr)
		if err := serv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	sig := <-sigChan
	fmt.Printf("Received signal: %v. Starting shutting down\n", sig)

	shuttingDownTime := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), shuttingDownTime)
	defer cancel()

	err = serv.Shutdown(ctx)
	time.Sleep(shuttingDownTime)

	if err == nil {
		log.Println("Server stopped gracefully")
	}
}
