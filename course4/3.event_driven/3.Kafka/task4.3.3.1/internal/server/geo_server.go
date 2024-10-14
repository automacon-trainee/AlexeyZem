package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"

	"metrics/internal/controller"
	"metrics/internal/service"
)

func NewGeoServer() (*http.Server, error) {
	geoService := service.NewGeodataService()
	redisClient := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	geoProxy := service.NewGeodataServiceProxy(geoService, redisClient)
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	responder := controller.NewResponder(logger)
	controll := controller.NewGeoController(responder, geoProxy)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	broker := os.Getenv("BROKER")
	rout := controller.NewGeoRouter(controll, broker)

	server := http.Server{
		Addr:         ":2080",
		Handler:      rout,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &server, nil
}
