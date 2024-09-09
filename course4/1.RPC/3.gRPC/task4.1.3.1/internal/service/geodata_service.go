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
	"github.com/prometheus/client_golang/prometheus"

	"metrics/internal/metrics"
	"metrics/internal/models"
)

type GeodataService interface {
	Search(geocode models.RequestAddressGeocode, res *models.ResponseAddress) error
	Geocode(address models.ResponseAddress, res *models.ResponseAddressGeocode) error
}

type GeodataServiceImpl struct {
	metrics *metrics.ProxyMetrics
}

func (s *GeodataServiceImpl) Search(geocode models.RequestAddressGeocode, res *models.ResponseAddress) error {
	metric := s.metrics.NewDurationHistogram("Search_to_api_histogram", "request duration Search in second in api",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metric.Observe(duration)
	}()
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f", geocode.Lat, geocode.Lng)
	body, err := ParseURLGet(url)
	if err != nil {
		*res = models.ResponseAddress{}
		return err
	}

	err = json.Unmarshal(body, res)
	if err != nil {
		return err
	}
	return nil
}
func (s *GeodataServiceImpl) Geocode(address models.ResponseAddress, res *models.ResponseAddressGeocode) error {
	metric := s.metrics.NewDurationHistogram("Geocode_to_api_histogram", "request duration Geocode in second in api",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metric.Observe(duration)
	}()
	q := GetQuery(address)
	request := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", q)

	body, err := ParseURLGet(request)
	if err != nil {
		return err
	}

	coord := []models.ResponseAddressGeocode{}
	err = json.Unmarshal(body, &coord)
	if err != nil {
		return err
	}
	if len(coord) == 0 {
		return errors.New("no result")
	}
	*res = coord[0]
	return nil
}

func NewGeodataService() *GeodataServiceImpl {
	return &GeodataServiceImpl{
		metrics: metrics.NewProxyMetrics(),
	}
}

type GeodataServiceProxy struct {
	service GeodataService
	client  *redis.Client
	metrics *metrics.ProxyMetrics
}

func (g *GeodataServiceProxy) Search(geocode models.RequestAddressGeocode, res *models.ResponseAddress) error {
	metric := g.metrics.NewDurationHistogram("Search_from_cache", "request duration Search in second in cache",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	start := time.Now()

	str, err := json.Marshal(geocode)
	if err != nil {
		return g.service.Search(geocode, res)
	}

	val, err := g.client.Get(string(str)).Result()
	if err != nil {
		errParse := g.service.Search(geocode, res)
		if errParse != nil {
			return errParse
		}
		if errors.Is(err, redis.Nil) {
			g.client.Set(string(str), *res, time.Hour)
		}
		return nil
	}

	err = json.Unmarshal([]byte(val), res)
	duration := time.Since(start).Seconds()
	metric.Observe(duration)
	return err
}

func (g *GeodataServiceProxy) Geocode(arg models.ResponseAddress, res *models.ResponseAddressGeocode) error {
	metric := g.metrics.NewDurationHistogram("Geocode_from_cache", "request duration Geocode in second in cache",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	start := time.Now()
	str, err := json.Marshal(arg)
	if err != nil {
		return g.service.Geocode(arg, res)
	}

	val, err := g.client.Get(string(str)).Result()
	if err != nil {
		errParse := g.service.Geocode(arg, res)
		if errParse != nil {
			return errParse
		}
		if errors.Is(err, redis.Nil) {
			g.client.Set(string(str), *res, time.Hour)
		}
		return nil
	}

	err = json.Unmarshal([]byte(val), res)
	duration := time.Since(start).Seconds()
	metric.Observe(duration)
	return err
}

func NewGeodataServiceProxy(serv GeodataService, client *redis.Client) *GeodataServiceProxy {
	return &GeodataServiceProxy{
		service: serv,
		client:  client,
		metrics: metrics.NewProxyMetrics(),
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
