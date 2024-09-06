package internal

import (
	"testing"
)

func TestNewServer(t *testing.T) {
	server := NewServer(NewRouter())
	if server == nil {
		t.Error("server is nil")
	}
}

func TestNewRouter(t *testing.T) {
	router := NewRouter()
	if router == nil {
		t.Error("router is nil")
	}
}
