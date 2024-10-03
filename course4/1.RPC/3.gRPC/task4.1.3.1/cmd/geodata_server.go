package main

import (
	"context"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"metrics/internal/API/gRPCGeo"
	"metrics/internal/models"
	"metrics/internal/service"
)

type GeoServiceServer struct {
	gRPCGeo.UnimplementedGeoServiceServer
	geoServer service.GeodataService
}

func (g *GeoServiceServer) SearchAnswer(_ context.Context, geocode *gRPCGeo.RequestAddressGeocode) (*gRPCGeo.ResponseAddress, error) {
	reqAddr := models.RequestAddressGeocode{Lat: geocode.Lat, Lng: geocode.Lng}
	res := models.ResponseAddress{}
	err := g.geoServer.Search(reqAddr, &res)
	response := gRPCGeo.ResponseAddress{
		Address: &gRPCGeo.Address{
			Country:     res.Address.Country,
			Road:        res.Address.Road,
			County:      res.Address.County,
			Town:        res.Address.Town,
			State:       res.Address.State,
			Postcode:    res.Address.Postcode,
			CountryCode: res.Address.CountryCode,
		},
	}
	return &response, err
}

func (g *GeoServiceServer) GeocodeAnswer(_ context.Context, address *gRPCGeo.Address) (*gRPCGeo.ResponseAddressGeocode, error) {
	reqAddress := models.ResponseAddress{
		Address: models.Address{
			Country:     address.Country,
			Road:        address.Road,
			County:      address.County,
			Town:        address.Town,
			State:       address.State,
			Postcode:    address.Postcode,
			CountryCode: address.CountryCode,
		},
	}
	res := models.ResponseAddressGeocode{}
	err := g.geoServer.Geocode(reqAddress, &res)
	response := gRPCGeo.ResponseAddressGeocode{
		Lat: res.Lat,
		Lon: res.Lon,
	}
	return &response, err
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	geoService := service.NewGeodataService()
	geoProxy := service.NewGeodataServiceProxy(geoService, redisClient)
	protocol := os.Getenv("RPC_PROTOCOL")
	err = rpc.Register(geoProxy)
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	switch protocol {
	case "rpc":
		startRPC(l)
	case "json-rpc":
		startJSONRPC(l)
	case "gRPC":
		startGRPC(l, geoProxy)
	}
}

func startRPC(l net.Listener) {
	log.Println("Listening on :1234 with protocol RPC")
	rpc.Accept(l)
}

func startJSONRPC(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}
		go jsonrpc.ServeConn(conn)
	}
}

func startGRPC(l net.Listener, geoProxy service.GeodataService) {
	server := grpc.NewServer()
	gRPCGeo.RegisterGeoServiceServer(server, &GeoServiceServer{
		geoServer: geoProxy,
	})
	log.Println("Listening on :1234 with protocol gRPC")
	if err := server.Serve(l); err != nil {
		log.Fatal(err)
	}
}
