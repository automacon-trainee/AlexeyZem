package controller

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"project/internal/API/gRPCAuth"
	mymiddleware "project/internal/middleware"
)

type ProfileController interface {
	GetProfile(w http.ResponseWriter, r *http.Request)
	TakeBook(w http.ResponseWriter, r *http.Request)
	ReturnBook(w http.ResponseWriter, r *http.Request)
}

func NewProfileRouter(controller ProfileController, auth gRPCAuth.AuthServiceClient) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	mid := mymiddleware.NewMiddleware(auth)
	r.Use(mid.AuthVerify)
	r.Get("/profile/{id}", controller.GetProfile)
	r.Post("/profile/take/{id}", controller.TakeBook)
	r.Post("/profile/return/{id}", controller.ReturnBook)
	return r
}
