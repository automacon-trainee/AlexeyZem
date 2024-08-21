package internal

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserStorage map[string]*User

var Storage = make(UserStorage)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	regReq := RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&regReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, ok := Storage[regReq.Username]; ok {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("User already exists"))
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(regReq.Password), bcrypt.MinCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Storage[regReq.Username] = &User{
		Username: regReq.Username,
		Password: string(hash),
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("User created"))
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

var (
	errBadUser = errors.New("invalid username or password")
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	logReq := LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&logReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := Storage[logReq.Username]; !ok {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(errBadUser.Error()))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(Storage[logReq.Username].Password), []byte(logReq.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	claims := jwt.MapClaims{
		"username": Storage[logReq.Username].Username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}
	_, token, _ := auth.Encode(claims)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(token))
}
