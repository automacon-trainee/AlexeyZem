package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"project/internal/library/models"
	"project/internal/myerror"
)

type LibraryService interface {
	GetAll(ctx context.Context) ([]models.Book, error)
	Create(ctx context.Context, book models.Book) error
	Take(ctx context.Context, id int) (models.Book, error)
	Return(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (models.Book, error)
}

type Responder interface {
	ErrorBadRequest(w http.ResponseWriter, err error)
	ErrorInternal(w http.ResponseWriter, err error)

	OutputJSON(w http.ResponseWriter, data any)
}

type Impl struct {
	service   LibraryService
	responder Responder
}

func NewController(service LibraryService, responder Responder) *Impl {
	return &Impl{
		service:   service,
		responder: responder,
	}
}

func (c *Impl) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	books, err := c.service.GetAll(ctx)
	if err != nil {
		c.responder.ErrorInternal(w, err)

		return
	}

	c.responder.OutputJSON(w, books)
}

func (c *Impl) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var book models.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		c.responder.ErrorBadRequest(w, err)

		return
	}

	if book.Count <= 0 {
		c.responder.ErrorBadRequest(w, errors.New("invalid count"))

		return
	}

	err = c.service.Create(ctx, book)
	if err != nil {
		c.responder.ErrorInternal(w, err)

		return
	}

	c.responder.OutputJSON(w, "Create Success")
}

func (c *Impl) Take(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		c.responder.ErrorBadRequest(w, err)

		return
	}
	book, err := c.service.Take(ctx, id)
	if err != nil {
		if errors.Is(err, myerror.ErrNotBook) {
			c.responder.ErrorBadRequest(w, err)

			return
		}
		c.responder.ErrorInternal(w, err)

		return
	}

	c.responder.OutputJSON(w, fmt.Sprintf("Take book %s", book.Title))
}

func (c *Impl) Return(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		c.responder.ErrorBadRequest(w, err)

		return
	}

	err = c.service.Return(ctx, id)
	if err != nil {
		if errors.Is(err, myerror.ErrNotBook) {
			c.responder.ErrorBadRequest(w, err)

			return
		}
		c.responder.ErrorInternal(w, err)

		return
	}

	c.responder.OutputJSON(w, "Return Success")
}
