package service

import (
	"context"
	"errors"
	"testing"

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
	serv := NewUserServiceImpl(MockRepo{})
	if serv == nil {
		t.Errorf("should not return nil")
	}
}

func TestUserServiceImpl_GetByEmail(t *testing.T) {
	serv := NewUserServiceImpl(MockRepo{})
	_, err := serv.GetUserByEmail("email")
	if err != nil {
		t.Errorf("should not return error")
	}
}

func TestUserServiceImpl_GetAll(t *testing.T) {
	serv := NewUserServiceImpl(MockRepo{})
	_, err := serv.GetAllUsers()
	if err != nil {
		t.Errorf("should not return error")
	}
}
