package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	ticker         = "/ticker"
	trades         = "/trades"
	orderBook      = "/orderBook"
	currency       = "/currency"
	candlesHistory = "/candles_history"
)

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

type Exmo struct {
	client *http.Client
	url    string
}

func (e *Exmo) GetBody(path string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", e.url+path, http.NoBody)
	if err != nil {
		return nil, err
	}
	resp, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return body, err
}

func (e *Exmo) GetTicker() (Ticker, error) {
	body, err := e.GetBody(ticker)
	if err != nil {
		return nil, err
	}
	var res Ticker

	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (e *Exmo) GetTrades(pairs ...string) (Trades, error) {
	body, err := e.GetBody(trades)
	if err != nil {
		return nil, err
	}
	var all Trades
	if err := json.Unmarshal(body, &all); err != nil {
		return nil, err
	}

	res := make(Trades)
	for _, pair := range pairs {
		res[pair] = all[pair]
	}
	return res, nil
}

func (e *Exmo) GetOrderBook(limit int, pairs ...string) (OrderBook, error) {
	body, err := e.GetBody(orderBook)
	if err != nil {
		return nil, err
	}
	var all OrderBook
	if err := json.Unmarshal(body, &all); err != nil {
		return nil, err
	}

	res := make(OrderBook)
	for i := 0; i < limit && i < len(pairs); i++ {
		res[pairs[i]] = all[pairs[i]]
	}
	return res, nil
}

func (e *Exmo) GetCurrencies() (Currencies, error) {
	body, err := e.GetBody(currency)
	if err != nil {
		return nil, err
	}
	var res Currencies
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (e *Exmo) GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error) {
	path := fmt.Sprintf("%s?symbol=%s&limit=%v&from=%v&to=%v", candlesHistory, pair, limit, start.Unix(), end.Unix())
	body, err := e.GetBody(path)
	if err != nil {
		return CandlesHistory{}, err
	}
	var res CandlesHistory
	if err := json.Unmarshal(body, &res); err != nil {
		return CandlesHistory{}, err
	}
	return res, nil
}

func (e *Exmo) GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error) {
	candles, err := e.GetCandlesHistory(pair, limit, start, end)
	if err != nil {
		return nil, err
	}
	res := make([]float64, len(candles.Candles))

	for i, candle := range candles.Candles {
		res[i] = candle.C
	}
	return res, nil
}

func NewExmo(opts ...func(exmo *Exmo)) *Exmo {
	exmo := &Exmo{
		client: http.DefaultClient,
		url:    "someURl",
	}
	for _, opt := range opts {
		opt(exmo)
	}
	return exmo
}

func WithClient(client *http.Client) func(exmo *Exmo) {
	return func(exmo *Exmo) {
		exmo.client = client
	}
}

func WithURL(url string) func(exmo *Exmo) {
	return func(exmo *Exmo) {
		exmo.url = url
	}
}

func main() {}
