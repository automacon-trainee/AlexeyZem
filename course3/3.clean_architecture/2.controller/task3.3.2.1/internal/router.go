package internal

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"

	"project/controller"
)

func NewRouter() *chi.Mux {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	userController := controller.NewUserController(controller.NewResponder(logger))
	geodataController := controller.NewGeodataController(controller.NewResponder(logger))
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Group(func(router chi.Router) {
		router.Use(jwtauth.Verifier(controller.Auth))
		router.Use(jwtauth.Authenticator)
		router.Post("/api/address/search", geodataController.Search)
		router.Post("/api/address/geocode", geodataController.Geocode)
	})
	router.Post("/api/login", userController.LoginUser)
	router.Post("/api/register", userController.CreateUser)
	router.NotFound(handlerNot)
	return router
}

func handlerNot(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte("not found"))
}
