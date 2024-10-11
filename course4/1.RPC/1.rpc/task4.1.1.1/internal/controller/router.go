package controller

import (
	"net/http"
	"net/http/pprof"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Controller interface {
	Register(w http.ResponseWriter, r *http.Request)
	Auth(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
	Geocode(w http.ResponseWriter, r *http.Request)
	GetByEmail(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
}

func NewRouter(controller Controller, auth *jwtauth.JWTAuth) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Group(func(router chi.Router) {
		router.Use(jwtauth.Verifier(auth))
		router.Use(jwtauth.Authenticator)
		router.Post("/api/address/search", controller.Search)
		router.Post("/api/address/geocode", controller.Geocode)
		router.Get("/api/users", controller.GetAllUsers)
		router.Get("/api/users/{email}", controller.GetByEmail)
		router.Get("/debug/pprof/", pprof.Index)
	})

	router.Get("/metrics", promhttp.Handler().ServeHTTP)
	router.Post("/api/users/login", controller.Auth)
	router.Post("/api/users/register", controller.Register)
	router.NotFound(handlerNot)

	return router
}

func handlerNot(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	str := []byte("not found")
	_, _ = w.Write(str)
}
