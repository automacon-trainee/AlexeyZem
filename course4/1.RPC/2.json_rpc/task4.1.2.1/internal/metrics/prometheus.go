package metrics

import (
	"log"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type metricsStorage map[string]any

func newCounter(name, help string) prometheus.Counter {
	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: name,
		Help: help,
	})
	if err := prometheus.Register(counter); err != nil {
		log.Println(err)
	}

	return counter
}

func newDurationHistogram(name, help string, buckets []float64) prometheus.Histogram {
	histogram := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    name,
		Help:    help,
		Buckets: buckets,
	})

	if err := prometheus.Register(histogram); err != nil {
		log.Println(err)
	}

	return histogram
}

type ProxyMetrics struct {
	store metricsStorage
}

func NewProxyMetrics() *ProxyMetrics {
	return &ProxyMetrics{
		store: make(metricsStorage),
	}
}

func (m *ProxyMetrics) NewCounter(name, help string) prometheus.Counter {
	val, ok := m.store[name+help]
	if !ok {
		counter := newCounter(name, help)
		m.store[name+help] = counter
		return counter
	}

	return val.(prometheus.Counter)
}

func (m *ProxyMetrics) NewDurationHistogram(name, help string, buckets []float64) prometheus.Histogram {
	bucketsStr := make([]string, len(buckets))
	for i, b := range buckets {
		bucketsStr[i] = strconv.FormatFloat(b, 'f', -1, 64)
	}
	buck := strings.Join(bucketsStr, ";")
	val, ok := m.store[name+help+buck]
	if !ok {
		histogram := newDurationHistogram(name, help, buckets)
		m.store[name+help+buck] = histogram
		return histogram
	}

	return val.(prometheus.Histogram)
}
