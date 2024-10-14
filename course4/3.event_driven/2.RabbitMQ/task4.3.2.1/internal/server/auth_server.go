package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"metrics/internal/API/gRPCAuth"
	"metrics/internal/controller"
	"metrics/internal/models"
	"metrics/internal/repository"
	"metrics/internal/service"
)

type AuthService interface {
	VerifyToken(token string) (*models.User, error)
}

type AuthServ struct {
	gRPCAuth.UnimplementedAuthServiceServer
	authServer AuthService
}

func (as *AuthServ) VerifyToken(ctx context.Context, token *gRPCAuth.Token) (*gRPCAuth.User, error) {
	user, err := as.authServer.VerifyToken(token.Token)
	if err != nil {
		return nil, err
	}
	res := &gRPCAuth.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Id:       user.ID,
	}

	return res, nil
}

func NewAuthServer() (*http.Server, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if err := godotenv.Load(); err != nil {
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

	secret := os.Getenv("SECRET_KEY")
	token := jwtauth.New("HS256", []byte(secret), nil)
	redisClient := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	authService := service.NewAuthServiceImpl(repo, token, redisClient)

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	responder := controller.NewResponder(logger)
	controll := controller.NewAuthController(responder, authService)
	rout := controller.NewAuthRouter(controll)

	server := http.Server{
		Addr:         ":3080",
		Handler:      rout,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go startGRPC(authService)

	return &server, nil
}

func startGRPC(auth AuthService) {
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("failed to listen gRPC: %v", err)
	}
	server := grpc.NewServer()
	gRPCAuth.RegisterAuthServiceServer(server, &AuthServ{
		authServer: auth,
	})

	log.Println("Listening on :1234 with protocol gRPC")

	if err := server.Serve(l); err != nil {
		log.Fatal(err)
	}
}
