package responder

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type Respond struct {
	logger *logrus.Logger
}

func NewResponder(logger *logrus.Logger) *Respond {
	return &Respond{logger: logger}
}

func (r *Respond) OutputJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json:charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&Response{Success: true, Data: data})
	if err != nil {
		r.logger.Error(err)
	}
	r.logger.Info(data)
}

func (r *Respond) ErrorUnAuthorized(w http.ResponseWriter, err error) {
	r.logger.Error(err)
	w.Header().Set("Content-Type", "application/json:charset=UTF-8")
	w.WriteHeader(http.StatusUnauthorized)
	errJ := json.NewEncoder(w).Encode(&Response{Success: false, Data: err.Error(), Message: "unauthorized"})
	if errJ != nil {
		r.logger.Error(err)
	}
}

func (r *Respond) ErrorBadRequest(w http.ResponseWriter, err error) {
	r.logger.Error(err)
	w.Header().Set("Content-Type", "application/json:charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	errJ := json.NewEncoder(w).Encode(&Response{Success: false, Data: err.Error(), Message: "bad request"})
	if errJ != nil {
		r.logger.Error(err)
	}
}

func (r *Respond) ErrorInternal(w http.ResponseWriter, err error) {
	r.logger.Error(err)
	w.Header().Set("Content-Type", "application/json:charset=UTF-8")
	if errors.Is(err, context.Canceled) {
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	errJ := json.NewEncoder(w).Encode(&Response{Success: false, Data: err.Error(), Message: "internal server error"})
	if errJ != nil {
		r.logger.Error(err)
	}
}

func (r *Respond) ErrorForbidden(w http.ResponseWriter, err error) {
	r.logger.Error(err)
	w.Header().Set("Content-Type", "application/json:charset=UTF-8")
	w.WriteHeader(http.StatusForbidden)
	errJ := json.NewEncoder(w).Encode(&Response{Success: false, Data: err.Error(), Message: "error forbidden"})
	if errJ != nil {
		r.logger.Error(err)
	}
}
