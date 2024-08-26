package controller

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(contr UserController) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Group(func(r chi.Router) {
		r.Get("/api/users", contr.GetAllUsers)
		r.Get("/api/books", contr.GetAllBooks)
		r.Get("/api/authors", contr.GetAllAuthors)
		r.Post("/api/books", contr.AddBook)
		r.Post("/api/take/{id}", contr.TakeBook)
		r.Post("/api/return", contr.ReturnBook)
		r.Get("/api/users/{id}", contr.GetUser)
		r.Get("/api/books/{id}", contr.GetBook)
		r.Post("/api/authors/{id}", contr.GetAuthor)
		r.NotFound(contr.NotFound)
	})
	return r
}
