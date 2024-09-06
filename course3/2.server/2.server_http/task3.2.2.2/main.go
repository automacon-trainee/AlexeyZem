// Контейнеризируй сервер, используя пакет net/http и Docker

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello world")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	log.Printf("Starting server on port %s", port)
	http.HandleFunc("/", handler)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
