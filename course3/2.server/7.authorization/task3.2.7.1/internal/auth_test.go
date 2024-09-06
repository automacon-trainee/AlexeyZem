package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginHandler(t *testing.T) {
	{
		user := User{Username: "testuser", Password: "testpassword"}
		hash, _ := bcrypt.GenerateFromPassword([]byte("testpassword"), bcrypt.MinCost)
		Storage["testuser"] = &User{
			Username: "testuser",
			Password: string(hash),
		}
		auth = jwtauth.New("HS256", []byte("secret"), nil)

		reqBody, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(reqBody)))
		w := httptest.NewRecorder()

		LoginHandler(w, req)
		body, _ := io.ReadAll(w.Body)

		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if string(body) == "" {
			t.Errorf("expected a token but got an empty string")
		}
	}
	{
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader("wrongLSON"))
		w := httptest.NewRecorder()
		LoginHandler(w, req)
		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	}
	{
		user := User{Username: "testuserWrong", Password: "testpassword"}
		reqBody, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(reqBody)))
		w := httptest.NewRecorder()

		LoginHandler(w, req)

		if status := w.Code; status != http.StatusUnauthorized {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
		}
	}
	{
		user := User{Username: "testuser", Password: "testpasswordwrong"}
		reqBody, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(reqBody)))
		w := httptest.NewRecorder()

		LoginHandler(w, req)

		if status := w.Code; status != http.StatusUnauthorized {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
		}
	}
}

func TestRegisterHandler(t *testing.T) {
	{
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader("wrongLSON"))
		w := httptest.NewRecorder()
		RegisterHandler(w, req)
		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	}
	{
		user := User{Username: "testuser", Password: "testpassword"}
		reqBody, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(string(reqBody)))
		w := httptest.NewRecorder()
		RegisterHandler(w, req)
		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
		body, _ := io.ReadAll(w.Body)
		if string(body) != "User already exists" {
			t.Errorf("handler not returned error")
		}
	}
	{
		user := User{Username: "Newtestuser", Password: "testpassword"}
		reqBody, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(string(reqBody)))
		w := httptest.NewRecorder()
		RegisterHandler(w, req)
		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
		body, _ := io.ReadAll(w.Body)
		if string(body) != "User created" {
			t.Errorf("handler not returned string")
		}
	}
}
