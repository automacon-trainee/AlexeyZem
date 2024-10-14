package controller

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"metrics/internal/API/gRPCAuth"
)

func GetProvider() gRPCAuth.AuthServiceClient {
	conn, err := grpc.NewClient("auth:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := gRPCAuth.NewAuthServiceClient(conn)

	return client
}
