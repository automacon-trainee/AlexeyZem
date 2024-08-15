package internal

import (
	"time"

	"github.com/cinar/indicator"
)

type Indicatorer interface {
	SMA(pair string, resolution, period int, from, to time.Time) ([]float64, error)
	EMA(pair string, resolution, period int, from, to time.Time) ([]float64, error)
}

type Indicator struct {
	exchange     Exchanger
	calculateSMA func(data []float64, period int) []float64
	calculateEMA func(data []float64, period int) []float64
}

func (i *Indicator) SMA(pair string, resolution, period int, from, to time.Time) ([]float64, error) {
	res, err := i.exchange.GetClosePrice(pair, resolution, from, to)
	if err != nil {
		return nil, err
	}
	res = i.calculateSMA(res, period)
	return res, nil
}

func (i *Indicator) EMA(pair string, resolution, period int, from, to time.Time) ([]float64, error) {
	res, err := i.exchange.GetClosePrice(pair, resolution, from, to)
	if err != nil {
		return nil, err
	}
	res = i.calculateEMA(res, period)
	return res, nil
}

type IndicatorOption func(*Indicator)

func WithSma(calculateSMA func(data []float64, period int) []float64) IndicatorOption {
	return func(indicator *Indicator) {
		indicator.calculateSMA = calculateSMA
	}
}

func WithEMA(calculateEMA func(data []float64, period int) []float64) IndicatorOption {
	return func(indicator *Indicator) {
		indicator.calculateEMA = calculateEMA
	}
}

func CalculateSma(closing []float64, period int) []float64 {
	return indicator.Sma(period, closing)
}

func CalculateEma(closing []float64, period int) []float64 {
	return indicator.Ema(period, closing)
}

func NewIndicator(exchange Exchanger, opts ...IndicatorOption) *Indicator {
	res := &Indicator{exchange: exchange}
	for _, opt := range opts {
		opt(res)
	}
	return res
}
