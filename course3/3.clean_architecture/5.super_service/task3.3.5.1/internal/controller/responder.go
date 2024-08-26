package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

type Responder interface {
	OutputJSON(w http.ResponseWriter, data any)

	ErrorBadRequest(w http.ResponseWriter, err error)
	ErrorInternalServerError(w http.ResponseWriter, err error)
	ErrorNotFound(w http.ResponseWriter)
}

type ResponderImpl struct {
	logger *log.Logger
}

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func (r *ResponderImpl) OutputJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&Response{Success: true, Data: data, Message: "success"})
	if err != nil {
		r.logger.Println(err)
		return
	}
}

func (r *ResponderImpl) ErrorBadRequest(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	errJ := json.NewEncoder(w).Encode(&Response{Success: false, Data: err.Error(), Message: "bad request"})
	if errJ != nil {
		r.logger.Println(errJ)
	}
}

func (r *ResponderImpl) ErrorInternalServerError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	errJ := json.NewEncoder(w).Encode(&Response{Success: false, Data: err.Error(), Message: "internal server error"})
	if errJ != nil {
		r.logger.Println(errJ)
	}
}

func (r ResponderImpl) ErrorNotFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	errJ := json.NewEncoder(w).Encode(&Response{Success: false, Data: "", Message: "not found"})
	if errJ != nil {
		r.logger.Println(errJ)
	}
}

func NewResponder(logger *log.Logger) Responder {
	return &ResponderImpl{logger: logger}
}
