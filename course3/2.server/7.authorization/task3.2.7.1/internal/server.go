package internal

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Group(func(router chi.Router) {
		router.Use(jwtauth.Verifier(auth))
		router.Use(jwtauth.Authenticator)
		router.Post("/api/address/search", handlerSearch)
		router.Post("/api/address/geocode", handlerGeocode)
	})
	router.Post("/api/login", LoginHandler)
	router.Post("/api/register", RegisterHandler)
	router.NotFound(handlerNot)
	return router
}

func handlerNot(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte("not found"))
}

type Server struct {
	router *chi.Mux
}

func NewServer(r *chi.Mux) *Server {
	return &Server{router: r}
}

func (s *Server) Start() error {
	return http.ListenAndServe(":8080", s.router)
}
