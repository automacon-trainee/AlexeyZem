package main

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	username := "Alex"
	role := "admin"
	email := "a@b.com"
	user := NewUser(1,
		WithUserEmail(email),
		WithUserRole(role),
		WithUserName(username),
	)
	if user.ID != 1 {
		t.Errorf("User.ID = %d, want 1", user.ID)
	}
	if user.Email != email {
		t.Errorf("User.Email = %s, want %s", user.Email, email)
	}
	if user.Role != role {
		t.Errorf("User.Role = %s, want %s", user.Role, role)
	}
	if user.Username != username {
		t.Errorf("User.Username = %s, want %s", user.Username, username)
	}
}
