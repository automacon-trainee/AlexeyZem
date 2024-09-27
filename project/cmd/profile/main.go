package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"project/internal/API/gRPCProfile"
	"project/internal/profile/controller"
	"project/internal/profile/models"
	"project/internal/profile/repository"
	"project/internal/profile/service"
	"project/internal/provider"
	"project/internal/responder"
)

type ProfileGRPC struct {
	gRPCProfile.UnimplementedProfileServiceServer
	profile controller.ProfileService
}

func (p *ProfileGRPC) Create(ctx context.Context, profile *gRPCProfile.Profile) (*gRPCProfile.Null, error) {
	err := p.profile.CreateProfile(ctx, models.Profile{ID: int(profile.Id), Name: profile.Name, Lastname: profile.LastName})
	if err != nil {
		return nil, err
	}

	return &gRPCProfile.Null{}, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " host=" + dbHost + " port=" + dbPort +
		" sslmode=" + dbSSLMode

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	retries := 10
	i := 1
	log.Printf("try to ping:%v", i)
	for err = db.Ping(); err != nil; err = db.Ping() {
		if i > retries {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
		i++
		log.Printf("try to ping:%v", i)
	}

	repo := repository.NewPostgresDB(db)
	err = repo.CreateTableProfile(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = repo.CreateProfileBookTable(ctx)
	if err != nil {
		log.Fatal(err)
	}

	serv := service.NewProfileServiceImpl(repo, provider.GetBookProvider())
	logger := logrus.New()
	resp := responder.NewResponder(logger)
	contr := controller.NewProfileController(serv, resp)
	router := controller.NewProfileRouter(contr, provider.GetAuthProvider())

	server := http.Server{
		Addr:         ":2000",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go StartGRPC(serv)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func StartGRPC(profile controller.ProfileService) {
	l, err := net.Listen("tcp", ":1236")
	if err != nil {
		log.Fatalf("failed to listen gRPC: %v", err)
	}
	server := grpc.NewServer()
	gRPCProfile.RegisterProfileServiceServer(server, &ProfileGRPC{
		profile: profile,
	})
	log.Println("Listening on :1236 with protocol gRPC")
	if err := server.Serve(l); err != nil {
		log.Fatal(err)
	}
}
