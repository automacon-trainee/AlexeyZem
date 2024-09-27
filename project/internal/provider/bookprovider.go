package provider

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"project/internal/API/gRPCBook"
)

func GetBookProvider() gRPCBook.BookServiceClient {
	conn, err := grpc.NewClient("library:1235", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	return gRPCBook.NewBookServiceClient(conn)
}
