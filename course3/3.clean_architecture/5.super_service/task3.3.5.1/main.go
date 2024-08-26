package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"golibrary/internal/controller"
	"golibrary/internal/repository"
	"golibrary/internal/service"
)

func main() {
	db := connectToDB()
	defer func() {
		_ = db.Close()
	}()
	repo, err := repository.NewLibraryRepo(db)
	if err != nil {
		panic(err)
	}
	serv, err := service.NewLibraryFacade(repo)
	if err != nil {
		panic(err)
	}
	cont := controller.NewController(serv)
	router := controller.NewRouter(cont)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	sig := <-sigChan
	fmt.Printf("Received signal: %v. Starting shutting down\n", sig)

	shuttingDownTime := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), shuttingDownTime)
	defer cancel()

	err = server.Shutdown(ctx)
	time.Sleep(shuttingDownTime)

	if err == nil {
		log.Println("Server stopped gracefully")
	}
}

func connectToDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	dataSource := "user=" + user + " password=" + password + " dbname=" + dbname + " host=" + host + " port=" + port + " sslmode=" + sslmode
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal("failed to connect:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("failed ping:", err)
	}
	return db
}
