package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"project/internal/API/gRPCAuth"
	"project/internal/auth/controller"
	"project/internal/auth/repository"
	"project/internal/auth/service"
	"project/internal/responder"
)

type AuthServ struct {
	gRPCAuth.UnimplementedAuthServiceServer
	authServer controller.AuthService
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

	repo := repository.NewPostgresDataBase(db)
	err = repo.CreateNewUserTable(ctx)
	if err != nil {
		log.Fatal(err)
	}

	secret := os.Getenv("SECRET_KEY")
	token := jwtauth.New("HS256", []byte(secret), nil)
	redisClient := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	authService := service.NewAuthServiceImpl(repo, token, redisClient)
	logger := logrus.New()
	resp := responder.NewResponder(logger)
	control := controller.NewAuthController(resp, authService)
	rout := controller.NewAuthRouter(control)
	server := http.Server{
		Addr:         ":1000",
		Handler:      rout,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go startGRPC(authService)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Fatal(err)
	}
}

func startGRPC(auth controller.AuthService) {
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
