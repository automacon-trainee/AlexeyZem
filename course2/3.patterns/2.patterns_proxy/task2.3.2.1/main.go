package main

import (
	"encoding/json"
	"fmt"

	"github.com/cinar/indicator"
)

type Indicator interface {
	StochPrice() ([]float64, []float64)
	RSI(period int) ([]float64, []float64)
	StochRSI(period int) ([]float64, []float64)
	SMA(period int) []float64
	MACD() ([]float64, []float64)
	EMA() []float64
}

func UnmarshalKLines(data []byte) (KLines, error) {
	var r KLines
	err := json.Unmarshal(data, &r)
	return r, err
}

type KLines struct {
	Pair    string   `json:"pair"`
	Candles []Candle `json:"candles"`
}

func (r *KLines) MarshalKLines() ([]byte, error) {
	return json.Marshal(r)
}

type Candle struct {
	T int64   `json:"t"`
	O float64 `json:"o"`
	C float64 `json:"c"`
	H float64 `json:"h"`
	L float64 `json:"l"`
	V float64 `json:"v"`
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

type LinesProxy struct {
	lines Indicator
	cache map[string][]float64
}

func (l *LinesProxy) StochPrice() (k, d []float64) {
	k, ok := l.cache["k_stochprice"]
	d, ok2 := l.cache["d_stochprice"]
	if !ok || !ok2 {
		k, d = l.lines.StochPrice()
		l.cache["k_stochprice"] = k
		l.cache["d_stochprice"] = d
	}
	return k, d
}

func (l *LinesProxy) RSI(period int) (rs, rsi []float64) {
	firstKey := fmt.Sprintf("rs_%v", period)
	secondKey := fmt.Sprintf("rsi_%v", period)
	rs, ok := l.cache[firstKey]
	rsi, ok2 := l.cache[secondKey]
	if !ok || !ok2 {
		rs, rsi = l.lines.RSI(period)
		l.cache[firstKey] = rs
		l.cache[secondKey] = rsi
	}
	return rs, rsi
}

func (l *LinesProxy) StochRSI(period int) (k, d []float64) {
	firstKey := fmt.Sprintf("k_stochrsi_%v", period)
	secondKey := fmt.Sprintf("d_stochrsi_%v", period)
	k, ok := l.cache[firstKey]
	d, ok2 := l.cache[secondKey]
	if !ok || !ok2 {
		k, d = l.lines.StochRSI(period)
		l.cache[firstKey] = k
		l.cache[secondKey] = d
	}
	return k, d
}

func (l *LinesProxy) SMA(period int) []float64 {
	key := fmt.Sprintf("sma_%v", period)
	sma, ok := l.cache[key]
	if !ok {
		sma = l.lines.SMA(period)
		l.cache[key] = sma
	}
	return sma
}

func (l *LinesProxy) MACD() (k, d []float64) {
	k, ok := l.cache["macd"]
	d, ok2 := l.cache["signal"]
	if !ok || !ok2 {
		k, d = l.lines.MACD()
		l.cache["macd"] = k
		l.cache["signal"] = d
	}
	return k, d
}

func (l *LinesProxy) EMA() []float64 {
	k, ok := l.cache["ema"]
	if !ok {
		k = l.lines.EMA()
		l.cache["ema"] = k
	}
	return k
}

func LoadKLines(data []byte) *Lines {
	kLines, err := UnmarshalKLines(data)
	if err != nil {
		panic(err)
	}
	t := &Lines{}
	for _, v := range kLines.Candles {
		t.closing = append(t.closing, v.C)
		t.high = append(t.high, v.H)
		t.low = append(t.low, v.L)
	}
	return t
}

func LoadKLinesProxy(data []byte) *LinesProxy {
	lines := LoadKLines(data)
	return &LinesProxy{
		lines: lines,
		cache: make(map[string][]float64),
	}
}

func main() {}
