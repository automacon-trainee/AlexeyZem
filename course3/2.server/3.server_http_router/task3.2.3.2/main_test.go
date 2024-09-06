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
	{
		req, _ := http.NewRequest("GET", "/group1/1", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("TestFirstHandler() returned %v", resp.Code)
		}
		if resp.Body.String() != "Group 1 Привет, мир 1" {
			t.Errorf("TestFirstHandler() returned %v", resp.Body.String())
		}
	}
	{
		req, _ := http.NewRequest("GET", "/group1/2", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("TestFirstHandler() returned %v", resp.Code)
		}
		if resp.Body.String() != "Group 1 Привет, мир 2" {
			t.Errorf("TestFirstHandler() returned %v", resp.Body.String())
		}
	}
	{
		req, _ := http.NewRequest("GET", "/group1/3", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("TestFirstHandler() returned %v", resp.Code)
		}
		if resp.Body.String() != "Group 1 Привет, мир 3" {
			t.Errorf("TestFirstHandler() returned %v", resp.Body.String())
		}
	}
}

func TestSecondHandler(t *testing.T) {
	r := NewRouter()
	{
		req, _ := http.NewRequest("GET", "/group2/1", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("TestFirstHandler() returned %v", resp.Code)
		}
		if resp.Body.String() != "Group 2 Привет, мир 1" {
			t.Errorf("TestFirstHandler() returned %v", resp.Body.String())
		}
	}
	{
		req, _ := http.NewRequest("GET", "/group2/2", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("TestFirstHandler() returned %v", resp.Code)
		}
		if resp.Body.String() != "Group 2 Привет, мир 2" {
			t.Errorf("TestFirstHandler() returned %v", resp.Body.String())
		}
	}
	{
		req, _ := http.NewRequest("GET", "/group2/3", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("TestFirstHandler() returned %v", resp.Code)
		}
		if resp.Body.String() != "Group 2 Привет, мир 3" {
			t.Errorf("TestFirstHandler() returned %v", resp.Body.String())
		}
	}
}

func TestThirdHandler(t *testing.T) {
	r := NewRouter()
	{
		req, _ := http.NewRequest("GET", "/group3/1", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("TestFirstHandler() returned %v", resp.Code)
		}
		if resp.Body.String() != "Group 3 Привет, мир 1" {
			t.Errorf("TestFirstHandler() returned %v", resp.Body.String())
		}
	}
	{
		req, _ := http.NewRequest("GET", "/group3/2", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("TestFirstHandler() returned %v", resp.Code)
		}
		if resp.Body.String() != "Group 3 Привет, мир 2" {
			t.Errorf("TestFirstHandler() returned %v", resp.Body.String())
		}
	}
	{
		req, _ := http.NewRequest("GET", "/group3/3", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("TestFirstHandler() returned %v", resp.Code)
		}
		if resp.Body.String() != "Group 3 Привет, мир 3" {
			t.Errorf("TestFirstHandler() returned %v", resp.Body.String())
		}
	}
}
