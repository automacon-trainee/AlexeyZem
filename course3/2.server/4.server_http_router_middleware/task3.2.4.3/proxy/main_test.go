package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	input string
	want  bool
}

func TestCheckUrl(t *testing.T) {
	tests := []TestCase{
		{input: "/", want: false},
		{input: "/api", want: true},
		{input: "/api/", want: true},
		{input: "/api/1397/fj", want: true},
		{input: "/some/api/fj", want: false},
	}
	for _, test := range tests {
		res := CheckURL(test.input)
		if res != test.want {
			t.Errorf("CheckUrl(%q) = %v; want %v", test.input, res, test.want)
		}
	}
}

func TestGetProxy(t *testing.T) {
	{
		res := GetProxy("http://localhost:1313")
		if res == nil {
			t.Errorf("GetProxy(\"http://localhost:1313\") = %v; want not nil", res)
		}
	}
	{
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("want panic")
			}
		}()
		GetProxy(string([]byte{0x7f}))
	}
}

func TestHandler(t *testing.T) {
	{
		req, _ := http.NewRequest("GET", "/api/", nil)
		resp := httptest.NewRecorder()
		handler(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", resp.Code, http.StatusOK)
		}
		if resp.Body.String() != "hello from API" {
			t.Errorf("handler returned unexpected body: got %v want %v", "hello from API", resp.Body.String())
		}
	}
	{
		req, _ := http.NewRequest("GET", "/1", nil)
		resp := httptest.NewRecorder()
		handler(resp, req)
		if resp.Code != http.StatusBadGateway {
			t.Errorf("handler returned wrong status code: got %v want %v", resp.Code, http.StatusOK)
		}
	}
}
