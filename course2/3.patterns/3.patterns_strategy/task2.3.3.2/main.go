package main

import (
	"time"
)

type GeneralIndicatorer interface {
	GetData(pair string, period int, from, to time.Time, indicator Indicatorer) ([]float64, error)
}
type GeneralGetData struct{}

func (g *GeneralGetData) GetData(pair string, period int, from, to time.Time, indicator Indicatorer) ([]float64, error) {
	return indicator.GetData(pair, period, from, to)
}

type Indicatorer interface {
	GetData(pair string, period int, from, to time.Time) ([]float64, error)
}

type SMA struct {
	exchanger Exchanger
	calculate func(data []float64, period int) []float64
}

func (s *SMA) GetData(pair string, period int, from, to time.Time) ([]float64, error) {
	data, err := s.exchanger.GetClosePrice(pair, 3, from, to)
	if err != nil {
		return nil, err
	}
	return s.calculate(data, period), nil
}

type EMA struct {
	exchanger Exchanger
	calculate func(data []float64, period int) []float64
}

func (e *EMA) GetData(pair string, period int, from, to time.Time) ([]float64, error) {
	data, err := e.exchanger.GetClosePrice(pair, 3, from, to)
	if err != nil {
		return nil, err
	}
	return e.calculate(data, period), nil
}

func calculateSMA(data []float64, period int) []float64 {
	if period <= 0 {
		return nil
	}
	var sma = make([]float64, len(data)/period)
	for i := range sma {
		sum := 0.0
		for _, d := range data[i*period : i*period+period] {
			sum += d
		}
		sma[i] = sum / float64(period)
	}
	return sma
}

func calculateEMA(data []float64, period int) []float64 {
	if len(data) == 0 || period <= 0 {
		return nil
	}
	alpha := 2.0 / (float64(period) + 1.0)
	ema := make([]float64, len(data))
	ema[0] = data[0]
	for i := 1; i < len(data); i++ {
		ema[i] = data[i]*alpha + (1-alpha)*ema[i-1]
	}
	return ema
}

type CandlesHistory struct {
	Candles []Candle `json:"candles"`
}
type Candle struct {
	T int64   `json:"t"`
	O float64 `json:"o"`
	C float64 `json:"c"`
	H float64 `json:"h"`
	L float64 `json:"l"`
	V float64 `json:"v"`
}

type Currencies []string

type OrderBook map[string]OrderBookPair

type OrderBookPair struct {
	AskQuantity string  `json:"ask_quantity"`
	AskAmount   string  `json:"ask_amount"`
	AskTop      string  `json:"ask_top"`
	BidQuantity string  `json:"bid_quantity"`
	BidAmount   string  `json:"bid_amount"`
	BidTop      string  `json:"bid_top"`
	Ask         [][]int `json:"ask"`
	Bid         [][]int `json:"bid"`
}

type Ticker map[string]TickerValue

type TickerValue struct {
	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
	LastTrade string `json:"last_trade"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Avg       string `json:"avg"`
	Vol       string `json:"vol"`
	VolCurr   string `json:"vol_curr"`
	Updated   int64  `json:"updated"`
}

type Trades map[string][]Pair

type Pair struct {
	TradeID  int    `json:"trade_id"`
	Type     string `json:"type"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
	Amount   string `json:"amount"`
	Date     int64  `json:"date"`
}

type Exchanger interface {
	GetTicker() (Ticker, error)
	GetTrades(pairs ...string) (Trades, error)
	GetOrderBook(limit int, pairs ...string) (OrderBook, error)
	GetCurrencies() (Currencies, error)
	GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error)
	GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error)
}

func main() {}
