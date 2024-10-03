package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"metrics/internal/metrics"
	"metrics/internal/models"
)

type AuthService interface {
	CreateUser(user models.User) error
	AuthUser(user models.User) (string, error)
	VerifyToken(token string) (*models.User, error)
}

type AuthResponder interface {
	OutputJSON(w http.ResponseWriter, data any)

	ErrorUnAuthorized(w http.ResponseWriter, err error)
	ErrorBadRequest(w http.ResponseWriter, err error)
	ErrorInternal(w http.ResponseWriter, err error)
}

type AuthControllerImpl struct {
	responder   AuthResponder
	serviceAuth AuthService
	metrics     *metrics.ProxyMetrics
}

func NewAuthController(responder AuthResponder, serviceAuth AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		responder:   responder,
		serviceAuth: serviceAuth,
		metrics:     metrics.NewProxyMetrics(),
	}
}

func (ac *AuthControllerImpl) Register(w http.ResponseWriter, r *http.Request) {
	histogram := ac.metrics.NewDurationHistogram("Register_endpoint_histogram", "time request to register endpoint",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	counter := ac.metrics.NewCounter("Register_endpoint_counter", "count request to register endpoint")
	counter.Inc()
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		histogram.Observe(duration)
	}()
	regReq := models.RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&regReq)
	if err != nil {
		ac.responder.ErrorBadRequest(w, err)
		return
	}
	err = ac.serviceAuth.CreateUser(models.User{Username: regReq.Username, Password: regReq.Password, Email: regReq.Email})
	if err != nil {
		ac.responder.ErrorInternal(w, err)
		return
	}
	ac.responder.OutputJSON(w, models.Data{Message: "user created"})
}

func (ac *AuthControllerImpl) Auth(w http.ResponseWriter, r *http.Request) {
	histogram := ac.metrics.NewDurationHistogram("Auth_endpoint_histogram", "time request to auth endpoint",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	counter := ac.metrics.NewCounter("Auth_endpoint_counter", "count request to auth endpoint")
	counter.Inc()
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		histogram.Observe(duration)
	}()
	logReq := models.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&logReq)
	if err != nil {
		ac.responder.ErrorBadRequest(w, err)
		return
	}
	token, err := ac.serviceAuth.AuthUser(models.User{Email: logReq.Email, Password: logReq.Password})
	if err != nil {
		ac.responder.ErrorUnAuthorized(w, err)
		return
	}
	ac.responder.OutputJSON(w, models.Data{Message: token})
}
