package controller

import (
	"net/http"
	"net/http/pprof"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type UserController interface {
	GetByEmail(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
}

func NewUserRouter(controller UserController) *chi.Mux {
	mid := NewMiddleware(GetProvider(), GetRabbitBroker())
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Group(func(router chi.Router) {
		router.Use(mid.AuthVerify)
		router.Get("/api/users", controller.GetAllUsers)
		router.Get("/api/users/{email}", controller.GetByEmail)
		router.Get("/debug/pprof/", pprof.Index)
		router.Get("/metrics", promhttp.Handler().ServeHTTP)
	})
	router.NotFound(handlerNot)
	return router
}

func handlerNot(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte("not found"))
}
