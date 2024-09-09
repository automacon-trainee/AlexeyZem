package controller

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(controller *ImplController) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Group(func(r chi.Router) {
		r.Get("/api/users", controller.AllUser)
		r.Get("/api/users/{username}", controller.SearchUser)
		r.Post("/api/users", controller.Create)
		r.Put("/api/users/{username}", controller.UpdateUser)
		r.Delete("/api/users/{username}", controller.DeleteUser)
	})
	return router
}
