// Разработай  веб-приложение  в  котором  реализован  слой  репозитория  с  использованием  PostgreSQL,
// следуя принципам чистой архитектуры, и объединить все в Docker-compose.
// Конфигурация базы данных должна обрабатываться с помощью godotenv.
// Все  компоненты  приложения,  включая  базу  данных,  следует  инкапсулировать  в
//  Docker-контейнерах  и
// управлять ими с помощью Docker-compose

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

	"projectrepo/internal/controller"
	"projectrepo/internal/service"
	"projectrepo/repository"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	db := ConnectToDB()
	resp := controller.NewResponder(logger)
	repo := repository.NewUserRepository(db)
	err := repo.CreateNewTable()
	if err != nil {
		logger.Fatal(err)
	}
	serv := service.NewService(repo)
	contr := controller.NewController(resp, serv)
	rout := controller.NewRouter(contr)
	server := http.Server{
		Addr:         ":8080",
		Handler:      rout,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(err)
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

func ConnectToDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dataSource := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " host=" + dbHost + " port=" + dbPort
	dataSource += " sslmode=disable"
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal(fmt.Errorf("error connecting to database: %w", err))
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(fmt.Errorf("error ping to database: %w", err))
	}
	return db
}
