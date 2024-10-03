package service

import (
	"context"
	"errors"
	"testing"

	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"

	"metrics/internal/models"
)

type MockRepo struct{}

func (m MockRepo) Create(ctx context.Context, user models.User) error {
	return nil
}

func (m MockRepo) GetByID(ctx context.Context, id string) (models.User, error) {
	return models.User{}, nil
}

func (m MockRepo) GetByEmail(ctx context.Context, email string) (models.User, error) {
	if email == "" {
		return models.User{}, errors.New("invalid email")
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	return models.User{Email: email, Password: string(password)}, nil
}

func (m MockRepo) List(ctx context.Context) ([]models.User, error) {
	return nil, nil
}

func TestNewUserServiceImpl(t *testing.T) {
	serv := NewUserServiceImpl(MockRepo{}, &jwtauth.JWTAuth{})
	if serv == nil {
		t.Errorf("should not return nil")
	}
}

func TestUserServiceImpl_Create(t *testing.T) {
	serv := NewUserServiceImpl(MockRepo{}, &jwtauth.JWTAuth{})
	err := serv.CreateUser(models.User{Email: "", Password: "some", Username: "username"})
	if err != nil {
		t.Errorf("should not return error")
	}
}

func TestUserServiceImpl_GetByEmail(t *testing.T) {
	serv := NewUserServiceImpl(MockRepo{}, &jwtauth.JWTAuth{})
	_, err := serv.GetUserByEmail("email")
	if err != nil {
		t.Errorf("should not return error")
	}
}

func TestUserServiceImpl_GetAll(t *testing.T) {
	serv := NewUserServiceImpl(MockRepo{}, &jwtauth.JWTAuth{})
	_, err := serv.GetAllUsers()
	if err != nil {
		t.Errorf("should not return error")
	}
}

func TestUserServiceImpl_AuthUser(t *testing.T) {
	serv := NewUserServiceImpl(MockRepo{}, &jwtauth.JWTAuth{})
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
