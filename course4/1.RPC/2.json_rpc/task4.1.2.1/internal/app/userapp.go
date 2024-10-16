package app

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

	"metrics/internal/server"
)

func RunUserApp() error {
	serv, err := server.NewServer()
	if err != nil {
		return err
	}

	errServer := make(chan error)
	go func() {
		log.Println("Server start with addr:", serv.Addr)
		if err := serv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errServer <- err
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	var sig os.Signal
	select {
	case err := <-errServer:
		return err
	case sig = <-sigChan:
	}
	fmt.Printf("Recieved signal: %v. Starting shutting down\n", sig)

	shuttingDownTime := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), shuttingDownTime)
	defer cancel()

	err = serv.Shutdown(ctx)
	time.Sleep(shuttingDownTime)

	if err == nil {
		log.Println("Server stopped gracefully")
	}

	return err
}
