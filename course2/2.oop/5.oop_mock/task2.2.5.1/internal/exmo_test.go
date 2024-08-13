package internal

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestGetTicker(t *testing.T) {
	{
		response := `{
  "BTC_USD": {
    "buy_price": "589.06",
    "sell_price": "592",
    "last_trade": "591.221",
    "high": "602.082",
    "low": "584.51011695",
    "avg": "591.14698808",
    "vol": "167.59763535",
    "vol_curr": "99095.17162071",
    "updated": 1470250973
  }
}`
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}))
		defer server.Close()
		exmo := NewExmo(WithURL(server.URL), WithClient(http.DefaultClient))
		data, err := exmo.GetTicker()
		if err != nil {
			t.Errorf("GetTicker() error: %v", err)
		}
		res := make(Ticker)
		res["BTC_USD"] = TickerValue{
			BuyPrice:  "589.06",
			SellPrice: "592",
			LastTrade: "591.221",
			High:      "602.082",
			Low:       "584.51011695",
			Avg:       "591.14698808",
			Vol:       "167.59763535",
			VolCurr:   "99095.17162071",
			Updated:   1470250973,
		}
		if !reflect.DeepEqual(data, res) {
			t.Errorf("GetTicker() = %v, want %v", data, res)
		}
	}

	{
		response := `{
  "BTC_USD: {
    "buy_price: "589.06",
    "sell_price": "592",
  }
}`
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}))
		exmo := NewExmo(WithURL(server.URL), WithClient(http.DefaultClient))
		_, err := exmo.GetTicker()
		if err == nil {
			t.Errorf("GetTicker() not error")
		}
	}

	{
		exmo := NewExmo(WithURL("wrong"), WithClient(http.DefaultClient))
		_, err := exmo.GetTicker()
		if err == nil {
			t.Errorf("GetTicker() not error")
		}
	}
}

func TestGetTrades(t *testing.T) {
	{
		response := `{trL:}`
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}))

		exmo := NewExmo(WithURL(server.URL), WithClient(http.DefaultClient))
		_, err := exmo.GetTrades()
		if err == nil {
			t.Errorf("GetTrades() error: %v", err)
		}
		defer server.Close()
	}

	{
		response := `{
  "BTC_USD": [
    {
      "trade_id": 3,
      "type": "sell",
      "price": "100",
      "quantity": "1",
      "amount": "100",
      "date": 1435488248
    }
  ]
}`
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}))
		defer server.Close()
		res := make(Trades)
		exmo := NewExmo(WithURL(server.URL), WithClient(http.DefaultClient))
		res["BTC_USD"] = []Pair{
			{TradeID: 3, Type: "sell", Price: "100", Quantity: "1", Amount: "100", Date: 1435488248},
		}
		data, err := exmo.GetTrades("BTC_USD")
		if err != nil {
			t.Errorf("GetTrades() error: %v", err)
		}
		if !reflect.DeepEqual(data, res) {
			t.Errorf("GetTrades() = %v, want %v", data, res)
		}
	}

	{
		exmo := NewExmo(WithURL("wrong"), WithClient(http.DefaultClient))
		_, err := exmo.GetTrades()
		if err == nil {
			t.Errorf("GetTrades() not error")
		}
	}
}

func TestGetOrderBook(t *testing.T) {
	{
		exmo := NewExmo(WithURL("wrong"), WithClient(http.DefaultClient))
		_, err := exmo.GetOrderBook(10)
		if err == nil {
			t.Errorf("GetOrderBook() not error")
		}
	}

	{
		response := `{
  "BTC_USD": {
    "ask_quantity": "3",
    "ask_amount": "500",
    "ask_top": "100",
    "bid_quantity": "1",
    "bid_amount": "99",
    "bid_top": "99",
    "ask": [
      [
        100,
        1,
        100
      ],
      [
        200,
        2,
        400
      ]
    ],
    "bid": [
      [
        99,
        1,
        99
      ]
    ]
  }
}`
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}))
		exmo := NewExmo(WithURL(server.URL), WithClient(http.DefaultClient))
		expexted := make(OrderBook)
		expexted["BTC_USD"] = OrderBookPair{AskQuantity: "3", AskAmount: "500", AskTop: "100", BidQuantity: "1", BidAmount: "99", BidTop: "99",
			Ask: [][]int{{100, 1, 100}, {200, 2, 400}}, Bid: [][]int{{99, 1, 99}}}
		res, err := exmo.GetOrderBook(1, "BTC_USD")
		if err != nil {
			t.Errorf("GetOrderBook() error: %v", err)
		}
		if !reflect.DeepEqual(res, expexted) {
			t.Errorf("GetOrderBook() = %v, want %v", res, expexted)
		}
	}

	{
		response := `{badJson:`
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}))
		exmo := NewExmo(WithURL(server.URL), WithClient(http.DefaultClient))
		_, err := exmo.GetOrderBook(1, "BTC_USD")
		if err == nil {
			t.Errorf("GetOrderBook() error: %v", err)
		}
	}
}

func TestCurrencies(t *testing.T) {
	{
		exmo := NewExmo(WithURL("wrong"), WithClient(http.DefaultClient))
		_, err := exmo.GetCurrencies()
		if err == nil {
			t.Errorf("GetCurrencies() not error")
		}
	}

	{
		response := `[
  "USD",
  "EUR",
  "BTC",
  "DOGE",
  "LTC"
]`
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}))
		exmo := NewExmo(WithURL(server.URL), WithClient(http.DefaultClient))
		expected := Currencies{"USD", "EUR", "BTC", "DOGE", "LTC"}
		data, err := exmo.GetCurrencies()
		if err != nil {
			t.Errorf("GetCurrencies() error: %v", err)
		}
		if !reflect.DeepEqual(data, expected) {
			t.Errorf("GetCurrencies() = %v, want %v", data, expected)
		}
	}

	{
		response := `[wrong Kson`
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}))
		exmo := NewExmo(WithURL(server.URL), WithClient(http.DefaultClient))
		_, err := exmo.GetCurrencies()
		if err == nil {
			t.Errorf("GetCurrencies() not error")
		}
	}
}

func TestGetHandlesHistory(t *testing.T) {
	{
		exmo := NewExmo(WithURL("wrong"), WithClient(http.DefaultClient))
		_, err := exmo.GetCandlesHistory("BTC_USD", 10, time.Now(), time.Now())
		if err == nil {
			t.Errorf("GetHandlesHistory() not error")
		}
	}

	{
		response := `{
  "candles": [
    {
      "t": 1585557000000,
      "o": 6590.6164,
      "c": 6602.3624,
      "h": 6618.78965693,
      "l": 6579.054,
      "v": 6.932754980000013
    }
  ]
}`
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}))
		exmo := NewExmo(WithURL(server.URL), WithClient(http.DefaultClient))
		res, err := exmo.GetCandlesHistory("BTC_USD", 10, time.Now(), time.Now())
		if err != nil {
			t.Errorf("GetHandlesHistory() error: %v", err)
		}
		expected := CandlesHistory{Candles: []Candle{{T: 1585557000000, O: 6590.6164, C: 6602.3624, H: 6618.78965693, L: 6579.054, V: 6.932754980000013}}}
		if !reflect.DeepEqual(expected, res) {
			t.Errorf("GetHandlesHistory() = %v, want %v", res, expected)
		}
	}

	{
		response := `{badJson:`
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}))
		exmo := NewExmo(WithURL(server.URL), WithClient(http.DefaultClient))
		_, err := exmo.GetCandlesHistory("fn", 10, time.Now(), time.Now())
		if err == nil {
			t.Errorf("GetHandlesHistory() not error")
		}
	}
}

func TestGetClosePrice(t *testing.T) {
	{
		exmo := NewExmo(WithURL("wrong"), WithClient(http.DefaultClient))
		_, err := exmo.GetClosePrice("BTC_USD", 10, time.Now(), time.Now())
		if err == nil {
			t.Errorf("GetHandlesHistory() not error")
		}
	}

	{
		response := `{badJson:`
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}))
		exmo := NewExmo(WithURL(server.URL), WithClient(http.DefaultClient))
		_, err := exmo.GetClosePrice("fn", 10, time.Now(), time.Now())
		if err == nil {
			t.Errorf("GetHandlesHistory() not error")
		}
	}

	{
		response := `{
  "candles": [
    {
      "t": 1585557000000,
      "o": 6590.6164,
      "c": 6602.3624,
      "h": 6618.78965693,
      "l": 6579.054,
      "v": 6.932754980000013
    }
  ]
}`
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}))
		exmo := NewExmo(WithURL(server.URL), WithClient(http.DefaultClient))
		res, err := exmo.GetClosePrice("BTC_USD", 10, time.Now(), time.Now())
		if err != nil {
			t.Errorf("GetHandlesHistory() error: %v", err)
		}
		expected := []float64{6602.3624}
		if !reflect.DeepEqual(expected, res) {
			t.Errorf("GetHandlesHistory() = %v, want %v", res, expected)
		}
	}

}
