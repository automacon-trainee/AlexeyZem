package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	responder Responder
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	regReq := RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&regReq)
	if err != nil {
		u.responder.ErrorBadRequest(w, err)
		return
	}
	if _, ok := Storage[regReq.Username]; ok {
		u.responder.OutputJSON(w, Data{Message: "user already exists"})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(regReq.Password), bcrypt.MinCost)
	if err != nil {
		u.responder.ErrorInternal(w, err)
		return
	}

	Storage[regReq.Username] = &User{
		Username: regReq.Username,
		Password: string(hash),
	}
	u.responder.OutputJSON(w, Data{Message: "user created"})
}

func (u *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	logReq := LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&logReq)
	if err != nil {
		u.responder.ErrorBadRequest(w, err)
		return
	}

	if _, ok := Storage[logReq.Username]; !ok {
		u.responder.ErrorUnAuthorized(w, errBadUser)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(Storage[logReq.Username].Password), []byte(logReq.Password))
	if err != nil {
		u.responder.ErrorUnAuthorized(w, errBadUser)
		return
	}

	claims := jwt.MapClaims{
		"username": Storage[logReq.Username].Username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}
	_, token, _ := Auth.Encode(claims)

	u.responder.OutputJSON(w, Data{Message: token})
}

func NewUserController(responder Responder) *UserController {
	return &UserController{
		responder: responder,
	}
}
