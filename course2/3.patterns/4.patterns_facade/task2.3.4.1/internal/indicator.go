package internal

type Indicatorer interface {
	SMA(period int) ([]float64, error)
	EMA(period int) ([]float64, error)
	LoadCandles(candles CandlesHistory)
}
