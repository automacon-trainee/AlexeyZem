package controller

import (
	"net/http"
	"testing"
)

type MockUserController struct{}

func (m *MockUserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserController) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserController) AddBook(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserController) AddUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserController) AddAuthor(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserController) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserController) TakeBook(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserController) ReturnBook(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserController) GetBook(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserController) GetUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserController) GetAuthor(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserController) NotFound(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func TestNewRouter(t *testing.T) {
	rt := NewRouter(&MockUserController{})
	if rt == nil {
		t.Errorf("NewRouter returned nil")
	}
}
