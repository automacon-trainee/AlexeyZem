package internal

import (
	"errors"
	"reflect"
	"time"
)

type TechnicFacade struct {
	line      *Lines
	exchanger Exchanger
}

func NewTechnicFacade() (*TechnicFacade, error) {
	facade := &TechnicFacade{}
	cur, err := facade.exchanger.GetCurrencies()
	if err != nil {
		return nil, err
	}
	trade, err := facade.exchanger.GetTrades(cur...)
	if err != nil {
		return nil, err
	}
	tickerVal, err := facade.exchanger.GetTicker()
	if err != nil {
		return nil, err
	}
	start := time.Unix(trade[cur[0]][0].Date, 0)
	end := time.Unix(tickerVal[cur[0]].Updated, 0)
	hist, err := facade.exchanger.GetCandlesHistory(cur[0], 10, start, end)
	if err != nil {
		return nil, err
	}
	facade.line = LoadLines(hist)
	return facade, nil
}

func (t *TechnicFacade) SMA(period int) ([]float64, error) {
	empty := Lines{}
	if reflect.DeepEqual(*t.line, empty) {
		return nil, errors.New("line is nil")
	}
	res := t.line.SMA(period)
	return res, nil
}

func (t *TechnicFacade) EMA(period int) ([]float64, error) {
	empty := Lines{}
	if reflect.DeepEqual(*t.line, empty) {
		return nil, errors.New("line is nil")
	}
	res := t.line.EMA(period)
	return res, nil
}

func (t *TechnicFacade) LoadCandles(candles CandlesHistory) {
	t.line = LoadLines(candles)
}
