package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"projectService/internal/model"
	"projectService/internal/service"

	"projectService/internal/custom_error"
)

type Controller interface {
	Register(w http.ResponseWriter, r *http.Request)
	Auth(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
	Geocode(w http.ResponseWriter, r *http.Request)
}

type GeoController struct {
	responder Responder
	service   service.GeoServicer
}

func NewGeoController(responder Responder, serv service.GeoServicer) *GeoController {
	return &GeoController{
		responder: responder,
		service:   serv,
	}
}

func (gc *GeoController) Register(w http.ResponseWriter, r *http.Request) {
	regReq := model.RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&regReq)
	if err != nil {
		gc.responder.ErrorBadRequest(w, err)
		return
	}
	err = gc.service.CreateUser(model.User{Username: regReq.Username, Password: regReq.Password})
	if err != nil && errors.Is(err, custom_error.ErrAlreadyExists) {
		gc.responder.ErrorUnAuthorized(w, err)
		return
	} else if err != nil {
		gc.responder.ErrorInternal(w, err)
		return
	}
	gc.responder.OutputJSON(w, model.Data{"user created"})
}

func (gc *GeoController) Auth(w http.ResponseWriter, r *http.Request) {
	logReq := model.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&logReq)
	if err != nil {
		gc.responder.ErrorBadRequest(w, err)
		return
	}
	token, err := gc.service.AuthUser(model.User{Username: logReq.Username, Password: logReq.Password})
	if err != nil {
		gc.responder.ErrorUnAuthorized(w, err)
		return
	}
	gc.responder.OutputJSON(w, model.Data{token})
}

func (gc *GeoController) Search(w http.ResponseWriter, r *http.Request) {
	var coord model.RequestAddressGeocode
	err := json.NewDecoder(r.Body).Decode(&coord)
	if err != nil {
		gc.responder.ErrorBadRequest(w, err)
		return
	}
	addres, err := gc.service.Search(coord)
	if err != nil {
		gc.responder.ErrorInternal(w, err)
		return
	}
	gc.responder.OutputJSON(w, addres)
}

func (gc *GeoController) Geocode(w http.ResponseWriter, r *http.Request) {
	var address model.ResponseAddress
	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		gc.responder.ErrorBadRequest(w, err)
		return
	}
	coord, err := gc.service.Geocode(address)
	if err != nil {
		gc.responder.ErrorInternal(w, err)
		return
	}
	gc.responder.OutputJSON(w, coord)
}
