package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"

	"projectrepo/internal"
	"projectrepo/internal/models"
	"projectrepo/internal/service"
)

type Controller interface {
	SearchUser(w http.ResponseWriter, r *http.Request)
	AllUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type ImplController struct {
	responder Responder
	service   service.UserService
}

func (c *ImplController) SearchUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	user, err := c.service.Get(username)
	if err != nil {
		c.checkError(w, err)
		return
	}
	c.responder.OutputJSON(w, user)
}

func (c *ImplController) AllUser(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")
	users, err := c.service.GetAll(limit, offset)
	if err != nil {
		c.responder.ErrorInternalServerError(w, err)
		return
	}
	c.responder.OutputJSON(w, users)
}

func (c *ImplController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	err := c.service.Delete(username)
	if err != nil {
		c.checkError(w, err)
		return
	}
	c.responder.OutputJSON(w, nil)
}

func (c *ImplController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		c.responder.ErrorBadRequest(w, err)
		return
	}
	err = c.service.Update(username, &newUser)
	if err != nil {
		c.checkError(w, err)
		return
	}
	c.responder.OutputJSON(w, newUser)
}

func (c *ImplController) Create(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		c.responder.ErrorBadRequest(w, err)
	}
	err = c.service.Create(&newUser)
	if err != nil {
		c.checkError(w, err)
		return
	}
	c.responder.OutputJSON(w, newUser)
}

func (c *ImplController) checkError(w http.ResponseWriter, err error) {
	if errors.Is(err, internal.BadRequestError) {
		c.responder.ErrorBadRequest(w, err)
	} else {
		c.responder.ErrorInternalServerError(w, err)
	}
}

func NewController(responder Responder, serv service.UserService) Controller {
	return &ImplController{
		responder: responder,
		service:   serv,
	}
}
