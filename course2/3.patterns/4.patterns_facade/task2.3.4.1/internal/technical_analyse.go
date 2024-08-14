package internal

import (
	"github.com/cinar/indicator"
)

type TechnicalAnalysis interface {
	StochPrice() ([]float64, []float64)
	RSI(period int) ([]float64, []float64)
	StochRSI(period int) ([]float64, []float64)
	SMA(period int) []float64
	MACD() ([]float64, []float64)
	EMA() []float64
}

type Lines struct {
	high    []float64
	low     []float64
	closing []float64
}

func (t *Lines) StochPrice() (k, d []float64) {
	k, d = indicator.StochasticOscillator(t.high, t.low, t.closing)
	return k, d
}

func (t *Lines) RSI(period int) (rs, rsi []float64) {
	rs, rsi = indicator.RsiPeriod(period, t.closing)
	return rs, rsi
}

func (t *Lines) StochRSI(period int) (k, d []float64) {
	_, rsi := t.RSI(period)
	k, d = indicator.StochasticOscillator(t.high, t.low, rsi)
	return k, d
}

func (t *Lines) SMA(period int) []float64 {
	return indicator.Sma(period, t.closing)
}
func (t *Lines) MACD() (_, _ []float64) {
	return indicator.Macd(t.closing)
}

func (t *Lines) EMA() []float64 {
	period := 5
	return indicator.Ema(period, t.closing)
}

func LoadLines(candles CandlesHistory) *Lines {
	t := &Lines{}
	for _, v := range candles.Candles {
		t.closing = append(t.closing, v.C)
		t.high = append(t.high, v.H)
		t.low = append(t.low, v.L)
	}
	return t
}
