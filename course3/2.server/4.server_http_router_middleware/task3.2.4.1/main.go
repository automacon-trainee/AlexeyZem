// Создай сервер на пакете go-chi. Сервер должен содержать несколько маршрутов с различными методами
// Все маршруты должны быть прологированы с помощью middleware-компонента logger

package main

import (
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
	_ = http.ListenAndServe(":8080", NewRouter())
}
