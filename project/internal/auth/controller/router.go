package controller

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	})
	router.NotFound(handlerNot)
	return router
}

func handlerNot(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("wrong path"))
}
