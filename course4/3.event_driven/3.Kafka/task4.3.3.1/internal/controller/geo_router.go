package controller

import (
	"net/http"
	"net/http/pprof"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type GeoController interface {
	Search(w http.ResponseWriter, r *http.Request)
	Geocode(w http.ResponseWriter, r *http.Request)
}

func NewGeoRouter(controller GeoController, broker string) *chi.Mux {
	mid := NewMiddleware(GetProvider(), GetBroker(broker))
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Group(func(router chi.Router) {
		router.Use(mid.AuthVerify)
		router.Use(mid.Limiter)
		router.Post("/api/address/search", controller.Search)
		router.Post("/api/address/geocode", controller.Geocode)
		router.Get("/debug/pprof/", pprof.Index)
		router.Get("/metrics", promhttp.Handler().ServeHTTP)
	})
	router.NotFound(handlerNot)
	return router
}
