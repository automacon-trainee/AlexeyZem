package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis"

	"pprof/internal/models"
)

type GeodataService interface {
	Search(geocode models.RequestAddressGeocode) (models.ResponseAddress, error)
	Geocode(address models.ResponseAddress) (models.ResponseAddressGeocode, error)
}
type GeodataServiceImpl struct{}

func (s *GeodataServiceImpl) Search(geocode models.RequestAddressGeocode) (models.ResponseAddress, error) {
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f", geocode.Lat, geocode.Lng)
	body, err := ParseURLGet(url)
	if err != nil {
		return models.ResponseAddress{}, err
	}

	address := &models.ResponseAddress{}
	err = json.Unmarshal(body, address)
	if err != nil {
		return models.ResponseAddress{}, err
	}
	return *address, nil
}
func (s *GeodataServiceImpl) Geocode(address models.ResponseAddress) (models.ResponseAddressGeocode, error) {
	q := GetQuery(address)
	request := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", q)

	body, err := ParseURLGet(request)
	if err != nil {
		return models.ResponseAddressGeocode{}, err
	}

	coord := []models.ResponseAddressGeocode{}
	err = json.Unmarshal(body, &coord)
	if err != nil {
		return models.ResponseAddressGeocode{}, err
	}
	return coord[0], nil
}

func NewGeodataService() GeodataService {
	return &GeodataServiceImpl{}
}

type GeodataServiceProxy struct {
	service GeodataService
	client  *redis.Client
}

func (g *GeodataServiceProxy) Search(geocode models.RequestAddressGeocode) (models.ResponseAddress, error) {
	str, err := json.Marshal(geocode)
	if err != nil {
		return g.service.Search(geocode)
	}

	val, err := g.client.Get(string(str)).Result()
	if err != nil {
		addr, errParse := g.service.Search(geocode)
		if errParse != nil {
			return models.ResponseAddress{}, errParse
		}
		if errors.Is(err, redis.Nil) {
			g.client.Set(string(str), addr, time.Hour)
		}
		return addr, nil
	}

	var res models.ResponseAddress
	err = json.Unmarshal([]byte(val), &res)
	return res, err
}

func (g *GeodataServiceProxy) Geocode(address models.ResponseAddress) (models.ResponseAddressGeocode, error) {
	str, err := json.Marshal(address)
	if err != nil {
		return g.service.Geocode(address)
	}

	val, err := g.client.Get(string(str)).Result()
	if err != nil {
		geo, errParse := g.service.Geocode(address)
		if errParse != nil {
			return models.ResponseAddressGeocode{}, errParse
		}
		if errors.Is(err, redis.Nil) {
			g.client.Set(string(str), geo, time.Hour)
		}
		return geo, nil
	}

	var res models.ResponseAddressGeocode
	err = json.Unmarshal([]byte(val), &res)
	return res, err
}

func NewGeodataServiceProxy(serv GeodataService, client *redis.Client) GeodataService {
	return &GeodataServiceProxy{
		service: serv,
		client:  client,
	}
}

func GetQuery(address models.ResponseAddress) string {
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
