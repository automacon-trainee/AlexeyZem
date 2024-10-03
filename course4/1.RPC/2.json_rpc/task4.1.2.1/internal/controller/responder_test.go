package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestNewResponder(t *testing.T) {
	resp := NewResponder(log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds))
	if resp == nil {
		t.Error("responder is nil")
	}
}

func TestErrorInternalServerError(t *testing.T) {
	resp := NewResponder(log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds))
	{
		errMy := errors.New("my error")
		w := httptest.NewRecorder()
		resp.ErrorInternal(w, errMy)
		if w.Code != http.StatusInternalServerError {
			t.Error("expected 500")
		}
		if w.Header().Get("Content-Type") != "application/json:charset=UTF-8" {
			t.Errorf("expected application/json; charset=UTF-8")
		}
		res := &Response{}
		err := json.NewDecoder(w.Body).Decode(res)
		if err != nil {
			t.Error(err)
		}
		want := Response{Success: false, Data: errMy.Error(), Message: "internal server error"}
		if *res != want {
			t.Error("response is wrong")
		}
	}
}

func TestErrorBadRequest(t *testing.T) {
	resp := NewResponder(log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds))
	{
		errMy := errors.New("my error")
		w := httptest.NewRecorder()
		resp.ErrorBadRequest(w, errMy)
		if w.Code != http.StatusBadRequest {
			t.Error("expected 400")
		}
		if w.Header().Get("Content-Type") != "application/json:charset=UTF-8" {
			t.Error("expected application/json; charset=UTF-8")
		}
		res := &Response{}
		err := json.NewDecoder(w.Body).Decode(res)
		if err != nil {
			t.Error(err)
		}
		want := Response{Success: false, Data: errMy.Error(), Message: "bad request"}
		if *res != want {
			t.Error("response is wrong")
		}
	}
}

func TestErrorForbidden(t *testing.T) {
	resp := NewResponder(log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds))
	{
		errMy := errors.New("my error")
		w := httptest.NewRecorder()
		resp.ErrorForbidden(w, errMy)
		if w.Code != http.StatusForbidden {
			t.Error("expected 403")
		}
		if w.Header().Get("Content-Type") != "application/json:charset=UTF-8" {
			t.Error("expected application/json; charset=UTF-8")
		}
		res := &Response{}
		err := json.NewDecoder(w.Body).Decode(res)
		if err != nil {
			t.Error(err)
		}
		want := Response{Success: false, Data: errMy.Error(), Message: "error forbidden"}
		if *res != want {
			t.Error("response is wrong")
		}
	}
}

func TestErrorUnAuthorized(t *testing.T) {
	resp := NewResponder(log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds))
	{
		errMy := errors.New("my error")
		w := httptest.NewRecorder()
		resp.ErrorUnAuthorized(w, errMy)
		if w.Code != http.StatusUnauthorized {
			t.Error("expected 401")
		}
		if w.Header().Get("Content-Type") != "application/json:charset=UTF-8" {
			t.Error("expected application/json; charset=UTF-8")
		}
		res := &Response{}
		err := json.NewDecoder(w.Body).Decode(res)
		if err != nil {
			t.Error(err)
		}
		want := Response{Success: false, Data: errMy.Error(), Message: "unauthorized"}
		if *res != want {
			t.Error("response is wrong")
		}
	}
}

type FakeData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestSuccess(t *testing.T) {
	resp := NewResponder(log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds))
	{
		data := FakeData{ID: 1, Name: "some"}
		w := httptest.NewRecorder()
		resp.OutputJSON(w, data)
		if w.Code != http.StatusOK {
			t.Error("expected 200")
		}
		if w.Header().Get("Content-Type") != "application/json:charset=UTF-8" {
			t.Error("expected application/json; charset=UTF-8")
		}
		res := &Response{}
		err := json.NewDecoder(w.Body).Decode(res)
		if err != nil {
			t.Error(err)
		}
	}
}
