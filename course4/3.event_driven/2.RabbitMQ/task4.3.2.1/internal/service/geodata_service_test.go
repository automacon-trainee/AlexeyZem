package service

import (
	"testing"

	"metrics/internal/models"
)

func TestParseUrlGet(t *testing.T) {
	_, err := ParseURLGet("some url")
	if err == nil {
		t.Errorf("should return error")
	}
	_, err = ParseURLGet("http://google.com")
	if err != nil {
		t.Errorf("should not return error")
	}
}

func TestGetQuery(t *testing.T) {
	address := models.Address{Road: "Road My", Town: "Town", State: "State"}
	query := GetQuery(models.ResponseAddress{Address: address})
	want := "Road+My+Town+State"
	if query != want {
		t.Errorf("want %s, got %s", want, query)
	}
}

func TestNewGeodataService(t *testing.T) {
	serv := NewGeodataService()
	if serv == nil {
		t.Errorf("should not return nil")
	}
}

func TestGeocode(t *testing.T) {
	serv := NewGeodataService()
	err := serv.Geocode(models.ResponseAddress{Address: models.Address{Road: "red square"}}, &models.ResponseAddressGeocode{})
	if err != nil {
		t.Errorf("should not return error")
	}
	err = serv.Geocode(models.ResponseAddress{}, &models.ResponseAddressGeocode{})
	if err == nil {
		t.Errorf("should return error")
	}
}

func TestSearch(t *testing.T) {
	serv := NewGeodataService()
	err := serv.Search(models.RequestAddressGeocode{}, &models.ResponseAddress{})
	if err != nil {
		t.Errorf("should not return error")
	}
}
