package controller

import (
	"errors"

	"github.com/go-chi/jwtauth"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserStorage map[string]*User

var Storage = make(UserStorage)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

var (
	errBadUser = errors.New("invalid username or password")
	Auth       = jwtauth.New("HS256", []byte("secret"), nil)
)

type Data struct {
	Message string `json:"message"`
}
