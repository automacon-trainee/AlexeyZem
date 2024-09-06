package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func handleRoute1(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello World"))
}

func handleRoute2(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello World 2"))
}

func handleRoute3(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello World 3"))
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/1", handleRoute1)
	r.Get("/2", handleRoute2)
	r.Get("/3", handleRoute3)
	return r
}

func main() {
	log.Fatal(http.ListenAndServe(":1313", NewRouter()))
}
