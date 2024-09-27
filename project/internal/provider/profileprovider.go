package provider

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"project/internal/API/gRPCProfile"
)

func GetProfileProvider() gRPCProfile.ProfileServiceClient {
	conn, err := grpc.NewClient("profile:1236", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	return gRPCProfile.NewProfileServiceClient(conn)
}
