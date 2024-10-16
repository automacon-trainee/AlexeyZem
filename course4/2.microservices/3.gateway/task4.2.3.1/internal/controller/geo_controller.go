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

type GeoResponder interface {
	OutputJSON(w http.ResponseWriter, data any)

	ErrorBadRequest(w http.ResponseWriter, err error)
	ErrorInternal(w http.ResponseWriter, err error)
}

type GeoControllerImpl struct {
	responder  GeoResponder
	serviceGeo GeodataService
	metrics    *metrics.ProxyMetrics
}

func NewGeoController(responder GeoResponder, servGeo GeodataService) *GeoControllerImpl {
	return &GeoControllerImpl{
		responder:  responder,
		serviceGeo: servGeo,
		metrics:    metrics.NewProxyMetrics(),
	}
}

func (gc *GeoControllerImpl) Search(w http.ResponseWriter, r *http.Request) {
	histogram := gc.metrics.NewDurationHistogram("Search_endpoint_histogram",
		"time request to search endpoint",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	counter := gc.metrics.NewCounter("Search_endpoint_counter", "count request to search endpoint")
	counter.Inc()
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		histogram.Observe(duration)
	}()

	var (
		coord   models.RequestAddressGeocode
		address models.ResponseAddress
	)

	if err := json.NewDecoder(r.Body).Decode(&coord); err != nil {
		gc.responder.ErrorBadRequest(w, err)
		return
	}

	if err := gc.serviceGeo.Search(coord, &address); err != nil {
		gc.responder.ErrorInternal(w, err)
		return
	}

	gc.responder.OutputJSON(w, address)
}

func (gc *GeoControllerImpl) Geocode(w http.ResponseWriter, r *http.Request) {
	histogram := gc.metrics.NewDurationHistogram("Geocode_endpoint_histogram",
		"time request to geocode endpoint",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	counter := gc.metrics.NewCounter("Geocode_endpoint_counter", "count request to geocode endpoint")
	counter.Inc()
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		histogram.Observe(duration)
	}()

	var (
		address models.ResponseAddress
		coord   models.ResponseAddressGeocode
	)

	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		gc.responder.ErrorBadRequest(w, err)
		return
	}

	if err := gc.serviceGeo.Geocode(address, &coord); err != nil {
		gc.responder.ErrorInternal(w, err)
		return
	}

	gc.responder.OutputJSON(w, coord)
}
