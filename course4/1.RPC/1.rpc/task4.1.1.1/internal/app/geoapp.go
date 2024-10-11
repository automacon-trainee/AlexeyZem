package app

import (
	"log"
	"net"
	"net/rpc"

	"github.com/go-redis/redis"

	"metrics/internal/service"
)

func RunGeoApp() error {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	geoService := service.NewGeodataService()
	geoProxy := service.NewGeodataServiceProxy(geoService, redisClient)
	err := rpc.Register(geoProxy)
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		return err
	}

	log.Println("Listening on :1234")
	rpc.Accept(l)

	return nil
}
