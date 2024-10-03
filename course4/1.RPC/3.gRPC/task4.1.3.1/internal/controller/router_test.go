package controller

import (
	"net/http"
	"testing"

	"github.com/go-chi/jwtauth"
)

type MockController struct{}

func (m MockController) Register(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m MockController) Auth(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m MockController) Search(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m MockController) Geocode(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m MockController) GetByEmail(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m MockController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func TestRouter(t *testing.T) {
	router := NewRouter(MockController{}, jwtauth.New("HS256", []byte("secretKey"), nil))
	if router == nil {
		t.Error("router is nil")
	}
}
