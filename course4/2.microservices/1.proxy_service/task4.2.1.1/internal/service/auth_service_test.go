package service

import (
	"testing"

	"github.com/go-chi/jwtauth"

	"metrics/internal/models"
)

func TestAuthServiceImpl_Create(t *testing.T) {
	serv := NewAuthServiceImpl(MockRepo{}, &jwtauth.JWTAuth{})
	err := serv.CreateUser(models.User{Email: "", Password: "some", Username: "username"})
	if err != nil {
		t.Errorf("should not return error")
	}
}

func TestAuthServiceImpl_AuthUser(t *testing.T) {
	serv := NewAuthServiceImpl(MockRepo{}, &jwtauth.JWTAuth{})
	_, err := serv.AuthUser(models.User{Email: ""})
	if err == nil {
		t.Errorf("should return error")
	}
	_, err = serv.AuthUser(models.User{Email: "email", Password: "email!"})
	if err == nil {
		t.Errorf("should return error")
	}
	_, err = serv.AuthUser(models.User{Email: "email", Password: "email"})
	if err != nil {
		t.Errorf("should not return error")
	}
}
