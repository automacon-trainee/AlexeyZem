package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"metrics/internal/metrics"
	"metrics/internal/models"
)

type GeodataService interface {
	Search(geocode models.RequestAddressGeocode, res *models.ResponseAddress) error
	Geocode(address models.ResponseAddress, res *models.ResponseAddressGeocode) error
}

type GeoControllerImpl struct {
	responder  Responder
	serviceGeo GeodataService
	metrics    *metrics.ProxyMetrics
}

func NewGeoController(responder Responder, servGeo GeodataService) *GeoControllerImpl {
	return &GeoControllerImpl{
		responder:  responder,
		serviceGeo: servGeo,
		metrics:    metrics.NewProxyMetrics(),
	}
}

func (gc *GeoControllerImpl) Search(w http.ResponseWriter, r *http.Request) {
	histogram := gc.metrics.NewDurationHistogram("Search_endpoint_histogram", "time request to search endpoint",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	counter := gc.metrics.NewCounter("Search_endpoint_counter", "count request to search endpoint")
	counter.Inc()
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		histogram.Observe(duration)
	}()
	var coord models.RequestAddressGeocode
	err := json.NewDecoder(r.Body).Decode(&coord)
	if err != nil {
		gc.responder.ErrorBadRequest(w, err)
		return
	}
	var address models.ResponseAddress
	err = gc.serviceGeo.Search(coord, &address)
	if err != nil {
		gc.responder.ErrorInternal(w, err)
		return
	}
	gc.responder.OutputJSON(w, address)
}

func (gc *GeoControllerImpl) Geocode(w http.ResponseWriter, r *http.Request) {
	histogram := gc.metrics.NewDurationHistogram("Geocode_endpoint_histogram", "time request to geocode endpoint",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	counter := gc.metrics.NewCounter("Geocode_endpoint_counter", "count request to geocode endpoint")
	counter.Inc()
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		histogram.Observe(duration)
	}()
	var address models.ResponseAddress
	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		gc.responder.ErrorBadRequest(w, err)
		return
	}
	var coord models.ResponseAddressGeocode
	err = gc.serviceGeo.Geocode(address, &coord)
	if err != nil {
		gc.responder.ErrorInternal(w, err)
		return
	}
	gc.responder.OutputJSON(w, coord)
}
