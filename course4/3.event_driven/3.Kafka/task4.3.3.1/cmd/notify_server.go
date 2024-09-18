package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"metrics/internal/API/gRPCNotify"

	"metrics/internal/broker_server"
)

type NotifyServ struct {
	gRPCNotify.UnimplementedNotifyServiceServer
}

func NewNotifyServ() *NotifyServ {
	return &NotifyServ{}
}

func (ns *NotifyServ) SendMessage(ctx context.Context, message *gRPCNotify.Message) (*gRPCNotify.Mock, error) {
	// имитация отправки уведомления на почту.
	log.Printf("send to %s, message:%s\n", message.Email, message.Text)
	return nil, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	broker := os.Getenv("BROKER")
	notifyServer := NewNotifyServ()
	go startGRPC(notifyServer)
	cl := GetNotifyGRPCClient()
	switch broker {
	case "kafka":
		broker_server.StartKafka(cl)
	case "rabbit":
		broker_server.StartRabbit(cl)
	default:
		log.Println("Unknown broker " + broker)
	}
	log.Println("Broker start sucessful")
}

func GetNotifyGRPCClient() gRPCNotify.NotifyServiceClient {
	conn, err := grpc.NewClient("notify:1235", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := gRPCNotify.NewNotifyServiceClient(conn)
	return client
}

func startGRPC(NotifyServ gRPCNotify.NotifyServiceServer) {
	l, err := net.Listen("tcp", ":1235")
	if err != nil {
		log.Fatalf("failed to listen gRPC: %v", err)
	}
	server := grpc.NewServer()
	gRPCNotify.RegisterNotifyServiceServer(server, NotifyServ)
	log.Println("Listening on :1235 with protocol gRPC")
	if err := server.Serve(l); err != nil {
		log.Fatal(err)
	}
}
