package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus"

	"metrics/internal/metrics"
	"metrics/internal/models"
)

type Responder interface {
	OutputJSON(w http.ResponseWriter, data any)

	ErrorUnAuthorized(w http.ResponseWriter, err error)
	ErrorBadRequest(w http.ResponseWriter, err error)
	ErrorInternal(w http.ResponseWriter, err error)
	ErrorForbidden(w http.ResponseWriter, err error)
}

type GeodataServiceRPC interface {
	Search(geocode models.RequestAddressGeocode) (models.ResponseAddress, error)
	Geocode(address models.ResponseAddress) (models.ResponseAddressGeocode, error)
}

type UserService interface {
	CreateUser(user models.User) error
	AuthUser(user models.User) (string, error)
	GetUserByEmail(email string) (models.User, error)
	GetAllUsers() ([]models.User, error)
}

type GeoController struct {
	responder   Responder
	serviceGeo  GeodataServiceRPC
	serviceUser UserService
	metrics     *metrics.ProxyMetrics
}

func NewGeoController(responder Responder, servUser UserService, servGeo GeodataServiceRPC) *GeoController {
	return &GeoController{
		responder:   responder,
		serviceGeo:  servGeo,
		serviceUser: servUser,
		metrics:     metrics.NewProxyMetrics(),
	}
}

func (gc *GeoController) Register(w http.ResponseWriter, r *http.Request) {
	histogram := gc.metrics.NewDurationHistogram("Register_endpoint_histogram",
		"time request to register endpoint",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	counter := gc.metrics.NewCounter("Register_endpoint_counter", "count request to register endpoint")
	counter.Inc()
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		histogram.Observe(duration)
	}()

	var regReq models.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&regReq)
	if err != nil {
		gc.responder.ErrorBadRequest(w, err)
		return
	}

	err = gc.serviceUser.CreateUser(models.User{Username: regReq.Username, Password: regReq.Password, Email: regReq.Email})
	if err != nil {
		gc.responder.ErrorInternal(w, err)
		return
	}

	gc.responder.OutputJSON(w, models.Data{Message: "user created"})
}

func (gc *GeoController) Auth(w http.ResponseWriter, r *http.Request) {
	histogram := gc.metrics.NewDurationHistogram("Auth_endpoint_histogram",
		"time request to auth endpoint",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	counter := gc.metrics.NewCounter("Auth_endpoint_counter", "count request to auth endpoint")
	counter.Inc()
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		histogram.Observe(duration)
	}()

	var logReq models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&logReq)
	if err != nil {
		gc.responder.ErrorBadRequest(w, err)
		return
	}

	token, err := gc.serviceUser.AuthUser(models.User{Email: logReq.Email, Password: logReq.Password})
	if err != nil {
		gc.responder.ErrorUnAuthorized(w, err)
		return
	}

	gc.responder.OutputJSON(w, models.Data{Message: token})
}

func (gc *GeoController) Search(w http.ResponseWriter, r *http.Request) {
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

	var coord models.RequestAddressGeocode

	err := json.NewDecoder(r.Body).Decode(&coord)
	if err != nil {
		gc.responder.ErrorBadRequest(w, err)
		return
	}

	address, err := gc.serviceGeo.Search(coord)
	if err != nil {
		gc.responder.ErrorInternal(w, err)
		return
	}

	gc.responder.OutputJSON(w, address)
}

func (gc *GeoController) Geocode(w http.ResponseWriter, r *http.Request) {
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

	var address models.ResponseAddress

	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		gc.responder.ErrorBadRequest(w, err)
		return
	}

	coord, err := gc.serviceGeo.Geocode(address)
	if err != nil {
		gc.responder.ErrorInternal(w, err)
		return
	}

	gc.responder.OutputJSON(w, coord)
}

func (gc *GeoController) GetByEmail(w http.ResponseWriter, r *http.Request) {
	histogram := gc.metrics.NewDurationHistogram("GetByEmail_endpoint_histogram",
		"time request to getByEmail endpoint",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	counter := gc.metrics.NewCounter("GetByEmail_endpoint_counter", "count request to getByEmail endpoint")
	counter.Inc()
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		histogram.Observe(duration)
	}()

	email := chi.URLParam(r, "email")
	user, err := gc.serviceUser.GetUserByEmail(email)
	if err != nil {
		gc.responder.ErrorInternal(w, err)
		return
	}

	gc.responder.OutputJSON(w, user)
}

func (gc *GeoController) GetAllUsers(w http.ResponseWriter, _ *http.Request) {
	histogram := gc.metrics.NewDurationHistogram("GetAllUser_endpoint_histogram",
		"time request to getAllUser endpoint",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	counter := gc.metrics.NewCounter("GetAllUser_endpoint_counter", "count request to getAllUser endpoint")
	counter.Inc()
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		histogram.Observe(duration)
	}()

	data, err := gc.serviceUser.GetAllUsers()
	if err != nil {
		gc.responder.ErrorInternal(w, err)
		return
	}

	gc.responder.OutputJSON(w, data)
}
