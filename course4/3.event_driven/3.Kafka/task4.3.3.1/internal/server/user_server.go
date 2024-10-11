package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"

	"metrics/internal/controller"
	"metrics/internal/repository"
	"metrics/internal/service"
)

func NewUserServer() (*http.Server, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " host=" + dbHost + " port=" + dbPort +
		" sslmode=" + dbSSLMode
	repo, err := repository.StartPostgressDataBase(ctx, connStr)
	if err != nil {
		return nil, err
	}

	userService := service.NewUserServiceImpl(repo)
	redisClient := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	userProxy := service.NewUserServiceProxy(userService, redisClient)
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	responder := controller.NewResponder(logger)

	controll := controller.NewUserController(responder, userProxy)

	broker := os.Getenv("BROKER")

	rout := controller.NewUserRouter(controll, broker)
	server := http.Server{
		Addr:         ":1080",
		Handler:      rout,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return &server, nil
}
