package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"pprof/internal/models"
	"pprof/internal/service"
)

type Controller interface {
	Register(w http.ResponseWriter, r *http.Request)
	Auth(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
	Geocode(w http.ResponseWriter, r *http.Request)
	GetByEmail(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
}

type GeoController struct {
	responder   Responder
	serviceGeo  service.GeodataService
	serviceUser service.UserService
}

func NewGeoController(responder Responder, servGeo service.GeodataService, servUser service.UserService) *GeoController {
	return &GeoController{
		responder:   responder,
		serviceGeo:  servGeo,
		serviceUser: servUser,
	}
}

func (gc *GeoController) Register(w http.ResponseWriter, r *http.Request) {
	regReq := models.RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&regReq)
	if err != nil {
		gc.responder.ErrorBadRequest(w, err)
		return
	}
	err = gc.serviceUser.CreateUser(models.User{Username: regReq.Username, Password: regReq.Password, Email: regReq.Email})
	if err != nil {
		gc.responder.ErrorInternal(w, err)
		return
	}
	gc.responder.OutputJSON(w, models.Data{Message: "user created"})
}

func (gc *GeoController) Auth(w http.ResponseWriter, r *http.Request) {
	logReq := models.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&logReq)
	if err != nil {
		gc.responder.ErrorBadRequest(w, err)
		return
	}
	token, err := gc.serviceUser.AuthUser(models.User{Email: logReq.Email, Password: logReq.Password})
	if err != nil {
		gc.responder.ErrorUnAuthorized(w, err)
		return
	}
	gc.responder.OutputJSON(w, models.Data{Message: token})
}

func (gc *GeoController) Search(w http.ResponseWriter, r *http.Request) {
	var coord models.RequestAddressGeocode
	err := json.NewDecoder(r.Body).Decode(&coord)
	if err != nil {
		gc.responder.ErrorBadRequest(w, err)
		return
	}
	address, err := gc.serviceGeo.Search(coord)
	if err != nil {
		gc.responder.ErrorInternal(w, err)
		return
	}
	gc.responder.OutputJSON(w, address)
}

func (gc *GeoController) Geocode(w http.ResponseWriter, r *http.Request) {
	var address models.ResponseAddress
	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		gc.responder.ErrorBadRequest(w, err)
		return
	}
	coord, err := gc.serviceGeo.Geocode(address)
	if err != nil {
		gc.responder.ErrorInternal(w, err)
		return
	}
	gc.responder.OutputJSON(w, coord)
}

func (gc *GeoController) GetByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	user, err := gc.serviceUser.GetUserByEmail(email)
	if err != nil {
		gc.responder.ErrorInternal(w, err)
	}
	gc.responder.OutputJSON(w, user)
}

func (gc *GeoController) GetAllUsers(w http.ResponseWriter, _ *http.Request) {
	data, err := gc.serviceUser.GetAllUsers()
	if err != nil {
		gc.responder.ErrorInternal(w, err)
	}
	gc.responder.OutputJSON(w, data)
}
