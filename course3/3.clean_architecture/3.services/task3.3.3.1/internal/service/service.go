package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"projectService/internal/controller"
	"projectService/internal/custom_error"
	"projectService/internal/model"
)

type Service struct{}

func NewService() controller.GeoServicer {
	return &Service{}
}

func (s *Service) CreateUser(user model.User) error {
	if _, ok := model.Storage[user.Username]; ok {
		return custom_error.ErrAlreadyExists
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	model.Storage[user.Username] = &model.User{
		Username: user.Username,
		Password: string(hash),
	}
	return nil
}

func (s *Service) AuthUser(user model.User) (string, error) {
	if _, ok := model.Storage[user.Username]; !ok {
		return "", custom_error.ErrBadUser
	}
	err := bcrypt.CompareHashAndPassword([]byte(model.Storage[user.Username].Password), []byte(user.Password))
	if err != nil {
		return "", custom_error.ErrBadUser
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}
	_, token, _ := model.Auth.Encode(claims)
	return token, nil
}

func (s *Service) Search(geocode model.RequestAddressGeocode) (model.ResponseAddress, error) {
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f", geocode.Lat, geocode.Lng)
	body, err := ParseURLGet(url)
	if err != nil {
		return model.ResponseAddress{}, err
	}

	address := &model.ResponseAddress{}
	err = json.Unmarshal(body, address)
	if err != nil {
		return model.ResponseAddress{}, err
	}
	return *address, nil
}
func (s *Service) Geocode(address model.ResponseAddress) (model.ResponseAddressGeocode, error) {
	q := GetQuery(address)
	request := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", q)

	body, err := ParseURLGet(request)
	if err != nil {
		return model.ResponseAddressGeocode{}, err
	}

	coord := []model.ResponseAddressGeocode{}
	err = json.Unmarshal(body, &coord)
	if err != nil {
		return model.ResponseAddressGeocode{}, err
	}
	return coord[0], nil
}

func GetQuery(address model.ResponseAddress) string {
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
