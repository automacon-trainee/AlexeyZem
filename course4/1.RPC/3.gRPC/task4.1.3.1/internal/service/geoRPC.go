package service

import (
	"net/rpc"

	"metrics/internal/models"
)

type GeoRPC struct {
	client *rpc.Client
}

func NewGeoRPC(client *rpc.Client) *GeoRPC {
	return &GeoRPC{
		client: client,
	}
}

func (g *GeoRPC) Search(geocode models.RequestAddressGeocode) (models.ResponseAddress, error) {
	var res models.ResponseAddress
	err := g.client.Call("GeodataServiceProxy.Search", geocode, &res)
	if err != nil {
		return models.ResponseAddress{}, err
	}
	return res, nil
}

func (g *GeoRPC) Geocode(address models.ResponseAddress) (models.ResponseAddressGeocode, error) {
	var res models.ResponseAddressGeocode
	err := g.client.Call("GeodataServiceProxy.Geocode", address, &res)
	if err != nil {
		return models.ResponseAddressGeocode{}, err
	}
	return res, nil
}
