package service

import (
	"context"

	"metrics/internal/API/gRPCGeo"
	"metrics/internal/models"
)

type GeoGRPC struct {
	client gRPCGeo.GeoServiceClient
}

func NewGeoGRPC(client gRPCGeo.GeoServiceClient) *GeoGRPC {
	return &GeoGRPC{
		client: client,
	}
}

func (g *GeoGRPC) Search(geocode models.RequestAddressGeocode) (models.ResponseAddress, error) {
	req := gRPCGeo.RequestAddressGeocode{Lat: geocode.Lat, Lng: geocode.Lng}
	resp, err := g.client.SearchAnswer(context.Background(), &req)
	if err != nil {
		return models.ResponseAddress{}, err
	}
	res := models.ResponseAddress{
		Address: models.Address{
			Road:        resp.Address.Road,
			Town:        resp.Address.Town,
			County:      resp.Address.County,
			State:       resp.Address.State,
			Postcode:    resp.Address.Postcode,
			Country:     resp.Address.Country,
			CountryCode: resp.Address.CountryCode,
		},
	}

	return res, nil
}

func (g *GeoGRPC) Geocode(address models.ResponseAddress) (models.ResponseAddressGeocode, error) {
	req := gRPCGeo.Address{
		Road:        address.Address.Road,
		Town:        address.Address.Town,
		County:      address.Address.County,
		State:       address.Address.State,
		Postcode:    address.Address.Postcode,
		Country:     address.Address.Country,
		CountryCode: address.Address.CountryCode,
	}

	resp, err := g.client.GeocodeAnswer(context.Background(), &req)
	if err != nil {
		return models.ResponseAddressGeocode{}, err
	}
	res := models.ResponseAddressGeocode{
		Lat: resp.Lat,
		Lon: resp.Lon,
	}

	return res, nil
}
