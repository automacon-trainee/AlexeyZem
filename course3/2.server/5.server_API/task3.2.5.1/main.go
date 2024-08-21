package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type ResponseAddress struct {
	Address Address `json:"address"`
}

type Address struct {
	Road        string `json:"road"`
	Town        string `json:"town"`
	County      string `json:"county"`
	State       string `json:"state"`
	Postcode    string `json:"postcode"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

type RequestAddressGeocode struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type ResponseAddressGeocode struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/api/address/search", handlerSearch)
	router.Post("/api/address/geocode", handlerGeocode)
	router.NotFound(handlerNot)
	return router
}

func handlerSearch(w http.ResponseWriter, r *http.Request) {
	var coord RequestAddressGeocode
	err := json.NewDecoder(r.Body).Decode(&coord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f", coord.Lat, coord.Lng)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, http.NoBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	address := &ResponseAddress{}
	err = json.Unmarshal(body, address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("You are in Country:%s, Town:%s, Road:%s", address.Address.Country, address.Address.Town, address.Address.Road)
	_, _ = w.Write([]byte(response))
}

func handlerGeocode(w http.ResponseWriter, r *http.Request) {
	var address ResponseAddress
	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	q := GetQuery(address)
	request := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", q)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", request, http.NoBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	coord := []ResponseAddressGeocode{}
	err = json.Unmarshal(body, &coord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Coordinate: Lon:%v, Lat:%v", coord[0].Lon, coord[0].Lat)
	_, _ = w.Write([]byte(response))
}

func GetQuery(address ResponseAddress) string {
	parts := []string{}
	parts = append(parts, strings.Split(address.Address.Road, " ")...)
	parts = append(parts, strings.Split(address.Address.Town, " ")...)
	parts = append(parts, strings.Split(address.Address.State, " ")...)
	parts = append(parts, strings.Split(address.Address.Country, " ")...)

	var sb strings.Builder
	for _, i := range parts {
		if i != "" {
			sb.WriteString("+")
			sb.WriteString(i)
		}
	}
	return strings.Trim(sb.String(), "+")
}

func handlerNot(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte("not found"))
}

func main() {
	_ = http.ListenAndServe(":8080", NewRouter())
}
