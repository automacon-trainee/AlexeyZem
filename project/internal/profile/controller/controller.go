package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"project/internal/profile/models"
)

type ID struct {
	ID int `json:"id"`
}

type ProfileService interface {
	GetProfile(ctx context.Context, id int) (*models.Profile, error)
	TakeBook(ctx context.Context, profileID, bookID int) error
	ReturnBook(ctx context.Context, profileID, bookID int) error
	CreateProfile(ctx context.Context, profile models.Profile) error
}

type Responder interface {
	ErrorBadRequest(w http.ResponseWriter, err error)
	ErrorInternal(w http.ResponseWriter, err error)

	OutputJSON(w http.ResponseWriter, data any)
}

type Impl struct {
	profileService ProfileService
	responder      Responder
}

func NewProfileController(profileService ProfileService, responder Responder) *Impl {
	return &Impl{
		profileService: profileService,
		responder:      responder,
	}
}

func (i *Impl) GetProfile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		i.responder.ErrorBadRequest(w, err)
		return
	}

	profile, err := i.profileService.GetProfile(r.Context(), idInt)
	if err != nil {
		i.responder.ErrorInternal(w, err)
		return
	}

	i.responder.OutputJSON(w, profile)
}

func (i *Impl) TakeBook(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(userID)
	if err != nil {
		i.responder.ErrorBadRequest(w, err)
		return
	}
	var bookID ID
	err = json.NewDecoder(r.Body).Decode(&bookID)
	if err != nil {
		i.responder.ErrorBadRequest(w, err)
		return
	}

	err = i.profileService.TakeBook(r.Context(), uID, bookID.ID)
	if err != nil {
		i.responder.ErrorInternal(w, err)
		return
	}

	i.responder.OutputJSON(w, "take book")
}

func (i *Impl) ReturnBook(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(userID)
	if err != nil {
		i.responder.ErrorBadRequest(w, err)
		return
	}
	var bookID ID
	err = json.NewDecoder(r.Body).Decode(&bookID)
	if err != nil {
		i.responder.ErrorBadRequest(w, err)
		return
	}

	err = i.profileService.ReturnBook(r.Context(), uID, bookID.ID)
	if err != nil {
		i.responder.ErrorInternal(w, err)
		return
	}

	i.responder.OutputJSON(w, "take book")
}
