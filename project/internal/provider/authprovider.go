package provider

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"project/internal/API/gRPCAuth"
)

func GetAuthProvider() gRPCAuth.AuthServiceClient {
	conn, err := grpc.NewClient("auth:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	return gRPCAuth.NewAuthServiceClient(conn)
}
