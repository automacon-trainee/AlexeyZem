package internal

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type mockExchanger struct{}

func (m mockExchanger) GetTicker() (Ticker, error) {
	panic("implement me")
}

func (m mockExchanger) GetTrades(pairs ...string) (Trades, error) {
	panic("implement me")
}

func (m mockExchanger) GetOrderBook(limit int, pairs ...string) (OrderBook, error) {
	panic("implement me")
}

func (m mockExchanger) GetCurrencies() (Currencies, error) {
	panic("implement me")
}

func (m mockExchanger) GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error) {
	panic("implement me")
}

func (m mockExchanger) GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error) {
	if end.Unix() < start.Unix() {
		return nil, fmt.Errorf("wrong input data")
	}
	res := []float64{100, 200, 300, 400, 500}
	return res, nil
}

func TestIndicator(t *testing.T) {
	mock := mockExchanger{}
	ind := NewIndicator(mock, WithSma(CalculateSma), WithEMA(CalculateEma))
	{
		res, err := ind.SMA("BTC_USD", 10, 10, time.Now().AddDate(0, 0, -1), time.Now())
		if err != nil {
			t.Errorf("not expected error: %v", err)
		}
		expected := []float64{100, 150, 200, 250, 300}
		if !reflect.DeepEqual(expected, res) {
			t.Errorf("want: %v, expected : %v", expected, res)
		}
	}

	{
		_, err := ind.SMA("BTC_USD", 10, 10, time.Now().AddDate(0, 0, 1), time.Now())
		if err == nil {
			t.Errorf("not error")
		}
	}

	{
		res, err := ind.EMA("BTC_USD", 10, 10, time.Now().AddDate(0, 0, -1), time.Now())
		if err != nil {
			t.Errorf("not expected error: %v", err)
		}
		expected := []float64{100, 118.18181818181819, 151.2396694214876, 196.46882043576258, 251.65630762926028}
		if !reflect.DeepEqual(expected, res) {
			t.Errorf("want: %v, expected : %v", expected, res)
		}
	}

	{
		_, err := ind.EMA("BTC_USD", 10, 10, time.Now().AddDate(0, 0, 1), time.Now())
		if err == nil {
			t.Errorf("not error")
		}
	}

}
