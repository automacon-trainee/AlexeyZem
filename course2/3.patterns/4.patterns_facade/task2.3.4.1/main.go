package main

import (
	"time"

	"facade/internal"
)

type Dashboarder interface {
	GetDashboard(pair string, opts ...func(dashboarder *Dashboard)) (DashboardData, error)
}

type DashboardData struct {
	Name           string
	CandlesHistory internal.CandlesHistory
	Indicators     map[string]IndicatorData
	period         int
	from           time.Time
	to             time.Time
}

type IndicatorData struct {
	Name     string
	Period   int
	Indicate []float64
}

type IndicatorOpt struct {
	Name      string
	Period    []int
	Indicator internal.Indicatorer
}

type Dashboard struct {
	exchange          internal.Exchanger
	withCandleHistory bool
	IndicatorOpts     []IndicatorOpt
	period            int
	from              time.Time
	to                time.Time
}

func (d *Dashboard) GetDashboard(pair string, opts ...func(dashboarder *Dashboard)) (DashboardData, error) {

	return DashboardData{}, nil
}

func NewDashboard(exchanger internal.Exchanger) *Dashboard {
	return &Dashboard{
		exchange: exchanger,
	}
}

func WithCandleHistory(period int, from, to time.Time) func(*Dashboard) {
	return func(d *Dashboard) {
		d.withCandleHistory = true
		d.period = period
		d.from = from
		d.to = to
	}
}

func main() {}
