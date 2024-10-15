package controller

import (
	"net/http"
	"net/http/pprof"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type AuthController interface {
	Register(w http.ResponseWriter, r *http.Request)
	Auth(w http.ResponseWriter, r *http.Request)
}

func NewAuthRouter(controller AuthController) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Group(func(router chi.Router) {
		router.Post("/api/auth/register", controller.Register)
		router.Post("/api/auth/login", controller.Auth)
		router.Get("/debug/pprof/", pprof.Index)
		router.Get("/metrics", promhttp.Handler().ServeHTTP)
	})
	router.NotFound(handlerNot)

	return router
}
