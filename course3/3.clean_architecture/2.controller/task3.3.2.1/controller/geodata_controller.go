package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Responder interface {
	OutputJSON(w http.ResponseWriter, data any)

	ErrorBadRequest(w http.ResponseWriter, err error)
	ErrorInternal(w http.ResponseWriter, err error)
}

type GeodataController struct {
	responder Responder
}

func (g *GeodataController) Search(w http.ResponseWriter, r *http.Request) {
	var coord RequestAddressGeocode
	err := json.NewDecoder(r.Body).Decode(&coord)
	if err != nil {
		g.responder.ErrorBadRequest(w, err)
		return
	}
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f", coord.Lat, coord.Lng)
	body, err := ParseURLGet(url)
	if err != nil {
		g.responder.ErrorInternal(w, err)
		return
	}

	address := &ResponseAddress{}
	err = json.Unmarshal(body, address)
	if err != nil {
		g.responder.ErrorInternal(w, err)
		return
	}
	g.responder.OutputJSON(w, address.Address)
}

func (g *GeodataController) Geocode(w http.ResponseWriter, r *http.Request) {
	var address ResponseAddress
	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		g.responder.ErrorBadRequest(w, err)
		return
	}

	q := GetQuery(address)
	request := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", q)

	body, err := ParseURLGet(request)
	if err != nil {
		g.responder.ErrorInternal(w, err)
		return
	}

	coord := []ResponseAddressGeocode{}
	err = json.Unmarshal(body, &coord)
	if err != nil {
		g.responder.ErrorInternal(w, err)
		return
	}
	g.responder.OutputJSON(w, ResponseAddressGeocode{Lon: coord[0].Lon, Lat: coord[0].Lat})
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

func ParseURLGet(url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, http.NoBody)
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func NewGeodataController(responder Responder) *GeodataController {
	return &GeodataController{
		responder: responder,
	}
}
