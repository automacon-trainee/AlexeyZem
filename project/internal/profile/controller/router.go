package controller

import (
	"net/http"
)

type ProfileController interface {
	GetProfile(w http.ResponseWriter, r *http.Request)
	TakeBook(w http.ResponseWriter, r *http.Request)
	ReturnBook(w http.ResponseWriter, r *http.Request)
}

type ProfileRouter struct {
	controller ProfileController
}

func NewProfileRouter(controller ProfileController) *ProfileRouter {
	return &ProfileRouter{
		controller: controller,
	}
}
