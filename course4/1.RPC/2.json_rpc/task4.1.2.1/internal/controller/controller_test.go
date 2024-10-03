package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"metrics/internal/models"
)

type MockServiceUser struct{}

func (m MockServiceUser) CreateUser(user models.User) error {
	if user.Username == "" {
		return errors.New("username is empty")
	}
	return nil
}

func (m MockServiceUser) AuthUser(user models.User) (string, error) {
	if user.Email == "" {
		return "", errors.New("email is empty")
	}
	return user.Username, nil
}

func (m MockServiceUser) GetUserByEmail(email string) (models.User, error) {
	if email == "" {
		return models.User{}, errors.New("email is empty")
	}
	return models.User{}, nil
}

func (m MockServiceUser) GetAllUsers() ([]models.User, error) {
	return []models.User{}, nil
}

type MockServiceGeoRPC struct{}

func (m MockServiceGeoRPC) Search(geocode models.RequestAddressGeocode) (models.ResponseAddress, error) {
	if geocode.Lat == 0 && geocode.Lng == 0 {
		return models.ResponseAddress{}, errors.New("geocode is empty")
	}
	return models.ResponseAddress{}, nil
}

func (m MockServiceGeoRPC) Geocode(address models.ResponseAddress) (models.ResponseAddressGeocode, error) {
	if address.Address.Road == "" {
		return models.ResponseAddressGeocode{}, errors.New("address is empty")
	}
	return models.ResponseAddressGeocode{}, nil
}

func TestNewController(t *testing.T) {
	resp := NewResponder(log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds))
	contr := NewGeoController(resp, MockServiceUser{}, MockServiceGeoRPC{})
	if contr == nil {
		t.Error("controller is nil")
	}
}

func TestRegister(t *testing.T) {
	resp := NewResponder(log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds))
	contr := NewGeoController(resp, MockServiceUser{}, MockServiceGeoRPC{})

	{
		body := "wrongJs}"
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/api/users/register", bytes.NewBufferString(body))
		contr.Register(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("wrong status code, want %d, got %d", http.StatusBadRequest, w.Code)
		}
	}
	{
		body, _ := json.Marshal(models.RegisterRequest{Username: "", Password: "", Email: ""})
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/api/users/register", bytes.NewBuffer(body))
		contr.Register(w, req)
		if w.Code != http.StatusInternalServerError {
			t.Errorf("wrong status code, want %d, got %d", http.StatusInternalServerError, w.Code)
		}
	}
	{
		body, _ := json.Marshal(models.RegisterRequest{Username: "Goland", Password: "go", Email: ""})
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/api/users/register", bytes.NewBuffer(body))
		contr.Register(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("wrong status code, want %d, got %d", http.StatusOK, w.Code)
		}
	}
}

func TestAuth(t *testing.T) {
	resp := NewResponder(log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds))
	contr := NewGeoController(resp, MockServiceUser{}, MockServiceGeoRPC{})

	{
		body := "wrongJs}"
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/api/users/login", bytes.NewBufferString(body))
		contr.Auth(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("wrong status code, want %d, got %d", http.StatusBadRequest, w.Code)
		}
	}
	{
		body, _ := json.Marshal(models.LoginRequest{Password: "", Email: ""})
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/api/users/login", bytes.NewBuffer(body))
		contr.Auth(w, req)
		if w.Code != http.StatusUnauthorized {
			t.Errorf("wrong status code, want %d, got %d", http.StatusUnauthorized, w.Code)
		}
	}
	{
		body, _ := json.Marshal(models.LoginRequest{Password: "go", Email: "Go@mail.ru"})
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/api/users/login", bytes.NewBuffer(body))
		contr.Auth(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("wrong status code, want %d, got %d", http.StatusOK, w.Code)
		}
	}
}

func TestGetByEmail(t *testing.T) {
	resp := NewResponder(log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds))
	contr := NewGeoController(resp, MockServiceUser{}, MockServiceGeoRPC{})
	{
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/api/users", nil)
		contr.GetByEmail(w, req)
		if w.Code != http.StatusInternalServerError {
			t.Errorf("wrong status code, want %d, got %d", http.StatusInternalServerError, w.Code)
		}
	}
}

func TestGetAllUser(t *testing.T) {
	resp := NewResponder(log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds))
	contr := NewGeoController(resp, MockServiceUser{}, MockServiceGeoRPC{})
	{
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/api/users", nil)
		contr.GetAllUsers(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("wrong status code, want %d, got %d", http.StatusOK, w.Code)
		}
	}
}

func TestGeocode(t *testing.T) {
	resp := NewResponder(log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds))
	contr := NewGeoController(resp, MockServiceUser{}, MockServiceGeoRPC{})

	{
		body := "wrongJs}"
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/api/address/geocode", bytes.NewBufferString(body))
		contr.Geocode(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("wrong status code, want %d, got %d", http.StatusBadRequest, w.Code)
		}
	}
	{
		body, _ := json.Marshal(models.ResponseAddress{Address: models.Address{}})
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/api/address/geocode", bytes.NewBuffer(body))
		contr.Geocode(w, req)
		if w.Code != http.StatusInternalServerError {
			t.Errorf("wrong status code, want %d, got %d", http.StatusInternalServerError, w.Code)
		}
	}
	{
		body, _ := json.Marshal(models.ResponseAddress{Address: models.Address{Road: "some road"}})
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/api/address/geocode", bytes.NewBuffer(body))
		contr.Geocode(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("wrong status code, want %d, got %d", http.StatusOK, w.Code)
		}
	}
}

func TestSearch(t *testing.T) {
	resp := NewResponder(log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds))
	contr := NewGeoController(resp, MockServiceUser{}, MockServiceGeoRPC{})

	{
		body := "wrongJs}"
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/api/address/Search", bytes.NewBufferString(body))
		contr.Search(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("wrong status code, want %d, got %d", http.StatusBadRequest, w.Code)
		}
	}
	{
		body, _ := json.Marshal(models.RequestAddressGeocode{Lng: 0, Lat: 0})
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/api/address/geocode", bytes.NewBuffer(body))
		contr.Search(w, req)
		if w.Code != http.StatusInternalServerError {
			t.Errorf("wrong status code, want %d, got %d", http.StatusInternalServerError, w.Code)
		}
	}
	{
		body, _ := json.Marshal(models.RequestAddressGeocode{Lng: 55, Lat: 12})
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/api/address/geocode", bytes.NewBuffer(body))
		contr.Search(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("wrong status code, want %d, got %d", http.StatusOK, w.Code)
		}
	}
}
