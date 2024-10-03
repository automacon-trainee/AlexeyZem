package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"golibrary/entities"
)

type MockServicer struct{}

func (m *MockServicer) StartService() error {
	//TODO implement me
	panic("implement me")
}

func (m *MockServicer) TakeBook(userID, bookID int) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockServicer) ReturnBook(book entities.Book) error {
	if book.ID == 0 {
		return errors.New("book id is 0")
	}
	return nil
}

func (m *MockServicer) AllUsersInfo() ([]entities.User, error) {
	return []entities.User{{ID: 1}}, nil
}

func (m *MockServicer) AllAuthorsInfo() ([]entities.Author, error) {
	return []entities.Author{{ID: 1}}, nil
}

func (m *MockServicer) AddBook(book entities.Book) error {
	if book.ID == 0 {
		return errors.New("book id should not be 0")
	}
	return nil
}

func (m *MockServicer) AddUser(user entities.User) error {
	if user.ID == 0 {
		return errors.New("user id should not be 0")
	}
	return nil
}

func (m *MockServicer) AddAuthor(author entities.Author) error {
	if author.ID == 0 {
		return errors.New("author id should not be 0")
	}
	return nil
}

func (m *MockServicer) GetAllBooks() ([]entities.Book, error) {
	return []entities.Book{{ID: 1}, {ID: 2}}, nil
}

func (m *MockServicer) BookInfo(bookID int) (entities.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockServicer) AuthorInfo(authorID int) (entities.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockServicer) UserInfo(userID int) (entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func TestNewController(t *testing.T) {
	contr := NewController(&MockServicer{})
	if contr == nil {
		t.Error("controller is nil")
	}
}

func TestGetUser(t *testing.T) {
	contr := NewController(&MockServicer{})
	{
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/api/users/{i}", nil)
		contr.GetUser(w, req)
		if w.Code != http.StatusBadRequest {
			t.Error("bad status")
		}
	}
}

func TestGetBook(t *testing.T) {
	contr := NewController(&MockServicer{})
	{
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/api/books/{i}", nil)
		contr.GetBook(w, req)
		if w.Code != http.StatusBadRequest {
			t.Error("bad status")
		}
	}
}

func TestGetAuthor(t *testing.T) {
	contr := NewController(&MockServicer{})
	{
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/api/authors/{i}", nil)
		contr.GetAuthor(w, req)
		if w.Code != http.StatusBadRequest {
			t.Error("bad status")
		}
	}
}

func TestGetAllUsers(t *testing.T) {
	contr := NewController(&MockServicer{})
	{
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/api/users", nil)
		contr.GetAllUsers(w, req)
		if w.Code != http.StatusOK {
			t.Error("bad status")
		}
	}
}

func TestGetAllAuthors(t *testing.T) {
	contr := NewController(&MockServicer{})
	{
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/api/authors", nil)
		contr.GetAllAuthors(w, req)
		if w.Code != http.StatusOK {
			t.Error("bad status")
		}
	}
}

func TestGetAllBooks(t *testing.T) {
	contr := NewController(&MockServicer{})
	{
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/api/books", nil)
		contr.GetAllBooks(w, req)
		if w.Code != http.StatusOK {
			t.Error("bad status")
		}
	}
}

func TestAddBook(t *testing.T) {
	contr := NewController(&MockServicer{})
	{
		w := httptest.NewRecorder()
		body, _ := json.Marshal(entities.Book{ID: 1, Name: "somebook"})
		req := httptest.NewRequest("POST", "/api/books", bytes.NewBuffer(body))
		contr.AddBook(w, req)
		if w.Code != http.StatusOK {
			t.Error("bad status")
		}
	}
	{
		w := httptest.NewRecorder()
		body := []byte("wrongJs")
		req := httptest.NewRequest("POST", "/api/books", bytes.NewBuffer(body))
		contr.AddBook(w, req)
		if w.Code != http.StatusBadRequest {
			t.Error("bad status")
		}
	}
	{
		w := httptest.NewRecorder()
		body, _ := json.Marshal(entities.Book{ID: 0, Name: "somebook"})
		req := httptest.NewRequest("POST", "/api/books", bytes.NewBuffer(body))
		contr.AddBook(w, req)
		if w.Code != http.StatusInternalServerError {
			t.Error("bad status")
		}
	}
}

func TestAddUser(t *testing.T) {
	contr := NewController(&MockServicer{})
	{
		w := httptest.NewRecorder()
		body, _ := json.Marshal(entities.User{ID: 1, Username: "someuser"})
		req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(body))
		contr.AddUser(w, req)
		if w.Code != http.StatusOK {
			t.Error("bad status")
		}
	}
	{
		w := httptest.NewRecorder()
		body := []byte("wrongJs")
		req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(body))
		contr.AddUser(w, req)
		if w.Code != http.StatusBadRequest {
			t.Error("bad status")
		}
	}
	{
		w := httptest.NewRecorder()
		body, _ := json.Marshal(entities.User{ID: 0, Username: "someuser"})
		req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(body))
		contr.AddUser(w, req)
		if w.Code != http.StatusInternalServerError {
			t.Error("bad status")
		}
	}
}

func TestAddAuthor(t *testing.T) {
	contr := NewController(&MockServicer{})
	{
		w := httptest.NewRecorder()
		body, _ := json.Marshal(entities.Author{ID: 1, Name: "someuser"})
		req := httptest.NewRequest("POST", "/api/authors", bytes.NewBuffer(body))
		contr.AddAuthor(w, req)
		if w.Code != http.StatusOK {
			t.Error("bad status")
		}
	}
	{
		w := httptest.NewRecorder()
		body := []byte("wrongJs")
		req := httptest.NewRequest("POST", "/api/authors", bytes.NewBuffer(body))
		contr.AddAuthor(w, req)
		if w.Code != http.StatusBadRequest {
			t.Error("bad status")
		}
	}
	{
		w := httptest.NewRecorder()
		body, _ := json.Marshal(entities.Author{ID: 0, Name: "somebook"})
		req := httptest.NewRequest("POST", "/api/authors", bytes.NewBuffer(body))
		contr.AddAuthor(w, req)
		if w.Code != http.StatusInternalServerError {
			t.Error("bad status")
		}
	}
}

func TestNotFound(t *testing.T) {
	contr := NewController(&MockServicer{})
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/", nil)
	contr.NotFound(w, req)
	if w.Code != http.StatusNotFound {
		t.Error("bad status")
	}
}

func TestTakeBook(t *testing.T) {
	contr := NewController(&MockServicer{})
	{
		w := httptest.NewRecorder()
		body, _ := json.Marshal(entities.Book{ID: 1, Name: "somebook"})
		req := httptest.NewRequest("POST", "/api/take/{ihov}", bytes.NewBuffer(body))
		contr.TakeBook(w, req)
		if w.Code != http.StatusBadRequest {
			t.Error("bad status")
		}
	}
	{
		w := httptest.NewRecorder()
		body := []byte("wrongJs")
		req := httptest.NewRequest("POST", "/api/take/10", bytes.NewBuffer(body))
		contr.TakeBook(w, req)
		if w.Code != http.StatusBadRequest {
			t.Error("bad status")
		}
	}
}

func TestReturnBook(t *testing.T) {
	contr := NewController(&MockServicer{})
	{
		w := httptest.NewRecorder()
		body, _ := json.Marshal(entities.Book{ID: 1, Name: "somebook"})
		req := httptest.NewRequest("POST", "/api/return", bytes.NewBuffer(body))
		contr.ReturnBook(w, req)
		if w.Code != http.StatusOK {
			t.Error("bad status")
		}
	}
	{
		w := httptest.NewRecorder()
		body, _ := json.Marshal(entities.Book{ID: 0, Name: "somebook"})
		req := httptest.NewRequest("POST", "/api/return", bytes.NewBuffer(body))
		contr.ReturnBook(w, req)
		if w.Code != http.StatusInternalServerError {
			t.Error("bad status")
		}
	}
	{
		w := httptest.NewRecorder()
		body := []byte("wrongJs")
		req := httptest.NewRequest("POST", "/api/return", bytes.NewBuffer(body))
		contr.ReturnBook(w, req)
		if w.Code != http.StatusBadRequest {
			t.Error("bad status")
		}
	}
}
