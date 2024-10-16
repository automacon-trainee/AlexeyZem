package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"

	"metrics/internal/controller"
	"metrics/internal/metrics"
	"metrics/internal/repository"
	"metrics/internal/service"
)

func NewServer() (*http.Server, error) {
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

	secretKey := os.Getenv("SECRET_KEY")
	token := jwtauth.New("HS256", []byte(secretKey), nil)
	userService := service.NewUserServiceImpl(repo, token)
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	userProxy := service.NewUserServiceProxy(userService, redisClient)
	geoService := service.NewGeodataService(metrics.NewProxyMetrics())
	geoProxy := service.NewGeodataServiceProxy(geoService, redisClient)
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	responder := controller.NewResponder(logger)
	controll := controller.NewGeoController(responder, geoProxy, userProxy)
	rout := controller.NewRouter(controll, token)
	server := http.Server{
		Addr:         ":8080",
		Handler:      rout,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return &server, nil
}
