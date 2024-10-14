package controller

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus"

	"metrics/internal/metrics"
	"metrics/internal/models"
)

type UserService interface {
	GetUserByEmail(email string) (models.User, error)
	GetAllUsers() ([]models.User, error)
}

type UserResponder interface {
	OutputJSON(w http.ResponseWriter, data any)

	ErrorBadRequest(w http.ResponseWriter, err error)
	ErrorInternal(w http.ResponseWriter, err error)
}

type UserControllerImpl struct {
	responder   UserResponder
	serviceUser UserService
	metrics     *metrics.ProxyMetrics
}

func NewUserController(responder UserResponder, servUser UserService) *UserControllerImpl {
	return &UserControllerImpl{
		responder:   responder,
		serviceUser: servUser,
		metrics:     metrics.NewProxyMetrics(),
	}
}

func (gc *UserControllerImpl) GetByEmail(w http.ResponseWriter, r *http.Request) {
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
	}

	gc.responder.OutputJSON(w, user)
}

func (gc *UserControllerImpl) GetAllUsers(w http.ResponseWriter, _ *http.Request) {
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
	}

	gc.responder.OutputJSON(w, data)
}
