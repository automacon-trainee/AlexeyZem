package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "metrics/cmd/gRPCGeo"
	"metrics/internal/controller"
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
	protocol := os.Getenv("RPC_PROTOCOL")

	geoProvider, err := GetGeoDataServiceRpc(protocol)
	if err != nil {
		return nil, err
	}
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	responder := controller.NewResponder(logger)
	controll := controller.NewGeoController(responder, userProxy, geoProvider)
	rout := controller.NewRouter(controll, token)
	server := http.Server{
		Addr:         ":8080",
		Handler:      rout,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return &server, nil
}

func GetGeoDataServiceRpc(protocol string) (controller.GeodataServiceRPC, error) {
	switch protocol {
	case "rpc":
		clientRPC, err := rpc.Dial("tcp", "geoservice:1234")
		if err != nil {
			return nil, err
		} else {
			return service.NewGeoRPC(clientRPC), nil
		}
	case "json-rpc":
		clientRPC, err := jsonrpc.Dial("tcp", "geoservice:1234")
		if err != nil {
			return nil, err
		} else {
			return service.NewGeoRPC(clientRPC), nil
		}
	case "gRPC":
		{
			conn, err := grpc.NewClient("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				return nil, err
			}
			grpcClient := pb.NewGeoServiceClient(conn)
			return service.NewGeoGRPC(grpcClient), nil
		}
	default:
		return nil, errors.New("invalid rpc protocol")
	}
}
