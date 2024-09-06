package internal

import (
	"github.com/go-chi/jwtauth"
)

var auth *jwtauth.JWTAuth

func init() {
	auth = jwtauth.New("HS256", []byte("secret"), nil)
}
