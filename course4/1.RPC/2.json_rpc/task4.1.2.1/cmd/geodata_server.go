package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"

	"metrics/internal/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	geoService := service.NewGeodataService()
	geoProxy := service.NewGeodataServiceProxy(geoService, redisClient)
	protocol := os.Getenv("RPC_PROTOCOL")
	err = rpc.Register(geoProxy)
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	switch protocol {
	case "rpc":
		startRPC(l)
	case "json-rpc":
		startJSONRPC(l)

	}
}

func startRPC(l net.Listener) {
	log.Println("Listening on :1234 with protocol RPC")
	rpc.Accept(l)
}

func startJSONRPC(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}
		go jsonrpc.ServeConn(conn)
	}
}
