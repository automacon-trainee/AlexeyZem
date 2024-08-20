package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRouter(t *testing.T) {
	r := NewRouter()
	if r == nil {
		t.Errorf("NewRouter() returned nil")
	}
}

func TestFirstHandler(t *testing.T) {
	r := NewRouter()
	req, _ := http.NewRequest("GET", "/1", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("TestFirstHandler() returned %v", resp.Code)
	}
	if resp.Body.String() != "Hello World" {
		t.Errorf("TestFirstHandler() returned %v", resp.Body.String())
	}
}

func TestSecondHandler(t *testing.T) {
	r := NewRouter()
	req, _ := http.NewRequest("GET", "/2", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("TestFirstHandler() returned %v", resp.Code)
	}
	if resp.Body.String() != "Hello World 2" {
		t.Errorf("TestFirstHandler() returned %v", resp.Body.String())
	}
}

func TestThirdHandler(t *testing.T) {
	r := NewRouter()
	req, _ := http.NewRequest("GET", "/3", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("TestFirstHandler() returned %v", resp.Code)
	}
	if resp.Body.String() != "Hello World 3" {
		t.Errorf("TestFirstHandler() returned %v", resp.Body.String())
	}
}
