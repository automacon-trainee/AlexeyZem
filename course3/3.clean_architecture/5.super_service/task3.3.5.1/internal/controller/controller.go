package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"golibrary/entities"
	"golibrary/internal/service"
)

type UserControllerImpl struct {
	Service   service.Servicer
	responder Responder
}

type UserController interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetAllAuthors(w http.ResponseWriter, r *http.Request)
	AddBook(w http.ResponseWriter, r *http.Request)
	GetAllBooks(w http.ResponseWriter, r *http.Request)
	TakeBook(w http.ResponseWriter, r *http.Request)
	ReturnBook(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetAuthor(w http.ResponseWriter, r *http.Request)
	NotFound(w http.ResponseWriter, r *http.Request)
}

func NewController(serv service.Servicer) UserController {
	return &UserControllerImpl{Service: serv}
}

func (us *UserControllerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		us.responder.ErrorBadRequest(w, err)
		return
	}
	user, err := us.Service.UserInfo(userID)
	if err != nil {
		us.responder.ErrorInternalServerError(w, err)
		return
	}
	us.responder.OutputJSON(w, user)
}

func (us *UserControllerImpl) GetAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		us.responder.ErrorBadRequest(w, err)
		return
	}
	user, err := us.Service.AuthorInfo(userID)
	if err != nil {
		us.responder.ErrorInternalServerError(w, err)
		return
	}
	us.responder.OutputJSON(w, user)
}

func (us *UserControllerImpl) GetBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		us.responder.ErrorBadRequest(w, err)
		return
	}
	book, err := us.Service.BookInfo(bookID)
	if err != nil {
		us.responder.ErrorInternalServerError(w, err)
		return
	}
	us.responder.OutputJSON(w, book)
}

func (us *UserControllerImpl) GetAllUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := us.Service.AllUsersInfo()
	if err != nil {
		us.responder.ErrorInternalServerError(w, err)
		return
	}
	us.responder.OutputJSON(w, users)
}

func (us *UserControllerImpl) GetAllAuthors(w http.ResponseWriter, _ *http.Request) {
	authors, err := us.Service.AllAuthorsInfo()
	if err != nil {
		us.responder.ErrorInternalServerError(w, err)
		return
	}
	us.responder.OutputJSON(w, authors)
}

func (us *UserControllerImpl) GetAllBooks(w http.ResponseWriter, _ *http.Request) {
	books, err := us.Service.GetAllBooks()
	if err != nil {
		us.responder.ErrorInternalServerError(w, err)
		return
	}
	us.responder.OutputJSON(w, books)
}

func (us *UserControllerImpl) AddBook(w http.ResponseWriter, r *http.Request) {
	var book entities.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		us.responder.ErrorBadRequest(w, err)
	}

	err = us.Service.AddBook(book)
	if err != nil {
		us.responder.ErrorInternalServerError(w, err)
		return
	}
	us.responder.OutputJSON(w, book)
}

func (us *UserControllerImpl) TakeBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var book entities.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		us.responder.ErrorBadRequest(w, err)
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		us.responder.ErrorBadRequest(w, err)
		return
	}

	err = us.Service.TakeBook(userID, book.ID)
	if err != nil {
		us.responder.ErrorInternalServerError(w, err)
		return
	}
	us.responder.OutputJSON(w, book)
}

func (us *UserControllerImpl) ReturnBook(w http.ResponseWriter, r *http.Request) {
	var book entities.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		us.responder.ErrorBadRequest(w, err)
	}

	err = us.Service.ReturnBook(book)
	if err != nil {
		us.responder.ErrorInternalServerError(w, err)
		return
	}
	us.responder.OutputJSON(w, book)
}

func (us *UserControllerImpl) NotFound(w http.ResponseWriter, _ *http.Request) {
	us.responder.ErrorNotFound(w)
}
