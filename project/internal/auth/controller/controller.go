package controller

import (
	"encoding/json"
	"net/http"

	"project/internal/auth/models"
)

type AuthService interface {
	CreateUser(user models.User) error
	AuthUser(user models.User) (string, error)
	VerifyToken(token string) (*models.User, error)
}

type Responder interface {
	ErrorBadRequest(w http.ResponseWriter, err error)
	ErrorInternal(w http.ResponseWriter, err error)
	ErrorUnAuthorized(w http.ResponseWriter, err error)

	OutputJSON(w http.ResponseWriter, data any)
}

type AuthControllerImpl struct {
	responder   Responder
	serviceAuth AuthService
}

func NewAuthController(responder Responder, serviceAuth AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		responder:   responder,
		serviceAuth: serviceAuth,
	}
}

func (ac *AuthControllerImpl) Register(w http.ResponseWriter, r *http.Request) {
	regReq := models.RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&regReq)
	if err != nil {
		ac.responder.ErrorBadRequest(w, err)
		return
	}
	err = ac.serviceAuth.CreateUser(models.User{Username: regReq.Username, Password: regReq.Password, Email: regReq.Email})
	if err != nil {
		ac.responder.ErrorInternal(w, err)
		return
	}
	ac.responder.OutputJSON(w, "user created")
}

func (ac *AuthControllerImpl) Auth(w http.ResponseWriter, r *http.Request) {
	logReq := models.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&logReq)
	if err != nil {
		ac.responder.ErrorBadRequest(w, err)
		return
	}
	token, err := ac.serviceAuth.AuthUser(models.User{Email: logReq.Email, Password: logReq.Password})
	if err != nil {
		ac.responder.ErrorUnAuthorized(w, err)
		return
	}
	ac.responder.OutputJSON(w, token)
}
