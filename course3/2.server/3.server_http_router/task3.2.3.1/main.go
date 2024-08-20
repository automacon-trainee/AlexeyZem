package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func handlerFirst(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello World"))
}

func handlerSecond(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello World 2"))
}

func handlerThird(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello World 3"))
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/1", handlerFirst)
	r.Get("/2", handlerSecond)
	r.Get("/3", handlerThird)
	r.NotFound(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Not Found"))
	})
	return r
}

func main() {
	r := NewRouter()
	_ = http.ListenAndServe(":8080", r)
}
