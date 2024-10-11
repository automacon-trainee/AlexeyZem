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
		r.Post("/api/users", contr.AddUser)
		r.Get("/api/users/{id}", contr.GetUser)
		r.Get("/api/books", contr.GetAllBooks)
		r.Post("/api/books", contr.AddBook)
		r.Get("/api/books/{id}", contr.GetBook)
		r.Get("/api/authors", contr.GetAllAuthors)
		r.Post("/api/authors", contr.AddAuthor)
		r.Post("/api/authors/{id}", contr.GetAuthor)
		r.Post("/api/take/{id}", contr.TakeBook)
		r.Post("/api/return", contr.ReturnBook)
		r.NotFound(contr.NotFound)
	})
	return r
}
