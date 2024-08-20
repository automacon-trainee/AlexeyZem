// Создай сервер на пакете go-chi.
// Сервер должен иметь несколько маршрутов с различными методами.
// Все маршруты должны быть зарегистрированы с использованием middleware для логирования
// с помощью zap logger.

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
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

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		start := time.Now()
		midle := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		defer func() {
			duration := time.Since(start)
			logger.Info(fmt.Sprintf("method: %s url:%s status:%d time:%s", r.Method, r.URL.Path, midle.Status(), duration))
		}()
		next.ServeHTTP(midle, r)
	})
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(LoggerMiddleware)
	r.Get("/1", handleRoute1)
	r.Get("/2", handleRoute2)
	r.Get("/3", handleRoute3)
	return r
}

func main() {
	http.ListenAndServe(":8080", NewRouter())
}
