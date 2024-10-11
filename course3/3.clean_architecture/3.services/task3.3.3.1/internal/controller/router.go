package controller

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"

	"projectService/internal/model"
)

func NewRouter(contr *GeoController) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Group(func(router chi.Router) {
		router.Use(jwtauth.Verifier(model.Auth))
		router.Use(jwtauth.Authenticator)
		router.Post("/api/address/search", contr.Search)
		router.Post("/api/address/geocode", contr.Geocode)
	})
	router.Post("/api/login", contr.Auth)
	router.Post("/api/register", contr.Register)
	router.NotFound(handlerNot)
	return router
}

func handlerNot(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte("not found"))
}
