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

	"project/internal/API/gRPCBook"
	"project/internal/provider"

	"project/internal/library/controller"
	"project/internal/library/repository"
	"project/internal/library/service"
	"project/internal/responder"
)

type LibGRPC struct {
	gRPCBook.UnimplementedBookServiceServer
	lib controller.LibraryService
}

func (l *LibGRPC) Take(ctx context.Context, id *gRPCBook.ID) (*gRPCBook.Book, error) {
	modelBook, err := l.lib.Take(ctx, int(id.Id))
	if err != nil {
		return nil, err
	}

	res := gRPCBook.Book{
		Id:     int64(modelBook.ID),
		Title:  modelBook.Title,
		Author: modelBook.Author,
		Count:  int64(modelBook.Count),
	}

	return &res, nil
}

func (l *LibGRPC) Return(ctx context.Context, id *gRPCBook.ID) (*gRPCBook.Book, error) {
	err := l.lib.Return(ctx, int(id.Id))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (l *LibGRPC) Get(ctx context.Context, id *gRPCBook.ID) (*gRPCBook.Book, error) {
	book, err := l.lib.Get(ctx, int(id.Id))
	if err != nil {
		return nil, err
	}

	res := gRPCBook.Book{
		Id:     int64(book.ID),
		Title:  book.Title,
		Author: book.Author,
		Count:  int64(book.Count),
	}

	return &res, nil
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
	err = repo.CreateNewBookTable(ctx)
	if err != nil {
		log.Fatal(err)
	}

	serv := service.NewLibraryServiceImpl(repo)
	logger := logrus.New()
	resp := responder.NewResponder(logger)
	contr := controller.NewController(serv, resp)
	router := controller.NewRouter(contr, provider.GetAuthProvider())

	server := http.Server{
		Addr:         "1001",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go StartGRPC(serv)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func StartGRPC(lib controller.LibraryService) {
	l, err := net.Listen("tcp", ":1235")
	if err != nil {
		log.Fatalf("failed to listen gRPC: %v", err)
	}
	server := grpc.NewServer()
	gRPCBook.RegisterBookServiceServer(server, &LibGRPC{
		lib: lib,
	})
	log.Println("Listening on :1235 with protocol gRPC")
	if err := server.Serve(l); err != nil {
		log.Fatal(err)
	}
}
