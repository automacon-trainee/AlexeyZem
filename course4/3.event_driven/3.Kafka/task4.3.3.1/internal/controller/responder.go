package controller

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type Respond struct {
	logger *log.Logger
}

func NewResponder(logger *log.Logger) *Respond {
	return &Respond{logger: logger}
}

func (r *Respond) OutputJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json:charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		r.logger.Println(err)
	}
}

func (r *Respond) ErrorUnAuthorized(w http.ResponseWriter, err error) {
	r.logger.Println(err)
	w.Header().Set("Content-Type", "application/json:charset=UTF-8")
	w.WriteHeader(http.StatusUnauthorized)
	if err := json.NewEncoder(w).Encode(&Response{Success: false, Data: err.Error(), Message: "unauthorized"}); err != nil {
		r.logger.Println(err)
	}
}

func (r *Respond) ErrorBadRequest(w http.ResponseWriter, err error) {
	r.logger.Println(err)
	w.Header().Set("Content-Type", "application/json:charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(&Response{Success: false, Data: err.Error(), Message: "bad request"}); err != nil {
		r.logger.Println(err)
	}
}

func (r *Respond) ErrorInternal(w http.ResponseWriter, err error) {
	r.logger.Println(err)
	w.Header().Set("Content-Type", "application/json:charset=UTF-8")
	if errors.Is(err, context.Canceled) {
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(&Response{Success: false, Data: err.Error(), Message: "internal server error"}); err != nil {
		r.logger.Println(err)
	}
}

func (r *Respond) ErrorForbidden(w http.ResponseWriter, err error) {
	r.logger.Println(err)
	w.Header().Set("Content-Type", "application/json:charset=UTF-8")
	w.WriteHeader(http.StatusForbidden)
	if err := json.NewEncoder(w).Encode(&Response{Success: false, Data: err.Error(), Message: "error forbidden"}); err != nil {
		r.logger.Println(err)
	}
}
