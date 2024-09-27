package controller

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"project/internal/API/gRPCAuth"
	mymiddleware "project/internal/middleware"
)

type Controller interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Take(w http.ResponseWriter, r *http.Request)
	Return(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

func NewRouter(controller Controller, auth gRPCAuth.AuthServiceClient) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	mid := mymiddleware.NewMiddleware(auth)
	r.Use(mid.AuthVerify)
	r.Post("/book/create", controller.Create)
	r.Get("/book/take/{id}", controller.Take)
	r.Get("/book/return/{id}", controller.Return)
	r.Get("/book/", controller.GetAll)

	return r
}
