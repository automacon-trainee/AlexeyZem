package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type TestCase struct {
	inputData   []float64
	inputPeriod int
	outputData  []float64
}

func TestCalculateSma(t *testing.T) {
	testCases := []TestCase{
		{[]float64{100.0, 200.0, 300.0}, 3, []float64{200}},
		{[]float64{100.0, 200.0, 300.0}, -10, nil},
	}
	for _, testCase := range testCases {
		res := calculateSMA(testCase.inputData, testCase.inputPeriod)
		if !reflect.DeepEqual(res, testCase.outputData) {
			t.Errorf("want: %v, result: %v", testCase.outputData, res)
		}
	}
}

func TestCalculateEMA(t *testing.T) {
	testCases := []TestCase{
		{[]float64{100.0, 200.0, 300.0}, 3, []float64{100, 150, 225}},
		{[]float64{100.0, 200.0, 300.0}, -10, nil},
	}
	for _, testCase := range testCases {
		res := calculateEMA(testCase.inputData, testCase.inputPeriod)
		if !reflect.DeepEqual(res, testCase.outputData) {
			t.Errorf("want: %v, result: %v", testCase.outputData, res)
		}
	}
}

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
func TestGeneralGetData(t *testing.T) {
	g := GeneralGetData{}
	mock := mockExchanger{}
	sma := SMA{calculate: calculateSMA, exchanger: mock}
	ema := EMA{calculate: calculateEMA, exchanger: mock}

	{
		data, err := g.GetData("some_pair", 3, time.Now().AddDate(0, 0, -1), time.Now(), &sma)
		if err != nil {
			t.Errorf("GetData error: %v", err)
		}
		want := []float64{200}
		if !reflect.DeepEqual(data, want) {
			t.Errorf("want: %v, got: %v", want, data)
		}
		_, err = g.GetData("some_pair", 3, time.Now().AddDate(0, 0, 1), time.Now(), &sma)
		if err == nil {
			t.Errorf("GetData error: %v", err)
		}
	}
	{
		data, err := g.GetData("some_pair", 3, time.Now().AddDate(0, 0, -1), time.Now(), &ema)
		if err != nil {
			t.Errorf("GetData error: %v", err)
		}
		want := []float64{100, 150, 225, 312.5, 406.25}
		if !reflect.DeepEqual(data, want) {
			t.Errorf("want: %v, got: %v", want, data)
		}
		_, err = g.GetData("some_pair", 3, time.Now().AddDate(0, 0, 1), time.Now(), &ema)
		if err == nil {
			t.Errorf("GetData error: %v", err)
		}
	}
}
