package main

const (
	ProductCocaCola = iota
	ProductPepsi
	ProductSprite
)

type Profitable interface {
	GetAverageProfitPercent() float64
	GetCurrentProfit() float64
	GetDifferenceProfit() float64
	GetAllData() []float64
	Average(prices []float64) float64
	Sum(prices []float64) float64
}

type Product struct {
	ProductID     int
	Sells         []float64
	Buys          []float64
	CurrentPrice  float64
	ProfitPercent float64
}

type StatisticProfit struct {
	product                 *Product
	getAverageProfit        func() float64
	getAverageProfitPercent func() float64
	getCurrentProfit        func() float64
	getDifferenceProfit     func() float64
	getAllData              func() []float64
}

func (s *StatisticProfit) GetAverageProfitPercent() float64 {
	return s.getAverageProfit()
}

func (s *StatisticProfit) GetCurrentProfit() float64 {
	return s.getCurrentProfit()
}

func (s *StatisticProfit) GetDifferenceProfit() float64 {
	return s.getDifferenceProfit()
}

func (s *StatisticProfit) GetAllData() []float64 {
	return s.getAllData()
}

func (s *StatisticProfit) Average(prices []float64) float64 {
	return s.Sum(prices) / float64(len(prices))
}

func (s *StatisticProfit) Sum(prices []float64) float64 {
	res := 0.0
	for _, price := range prices {
		res += price
	}
	return res
}

func (s *StatisticProfit) SetProduct(product *Product) {
	s.product = product
}

func NewStatisticProfit(opts ...func(*StatisticProfit)) Profitable {
	p := &StatisticProfit{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func WithAverageProfit(s *StatisticProfit) {
	s.getAverageProfit = func() float64 {
		return s.Average(s.product.Sells) - s.Average(s.product.Buys)
	}
}

func WithAverageProfitPercent(s *StatisticProfit) {
	s.getAverageProfitPercent = func() float64 {
		return 100 * (s.Average(s.product.Sells) - s.Average(s.product.Buys)) / s.Average(s.product.Buys)
	}
}

func WithCurrentProfit(s *StatisticProfit) {
	s.getCurrentProfit = func() float64 {
		return s.product.CurrentPrice - s.product.CurrentPrice*(100-s.product.ProfitPercent)/100
	}
}

func WithDifferenceProfit(s *StatisticProfit) {
	s.getDifferenceProfit = func() float64 {
		return s.product.CurrentPrice - s.Average(s.product.Sells)
	}
}

func WithAllData(s *StatisticProfit) {
	s.getAllData = func() []float64 {
		res := make([]float64, 0, 4)
		if s.getAverageProfit != nil {
			res = append(res, s.getAverageProfit())
		}
		if s.getAverageProfitPercent != nil {
			res = append(res, s.getAverageProfitPercent())
		}
		if s.getCurrentProfit != nil {
			res = append(res, s.getCurrentProfit())
		}
		if s.getDifferenceProfit != nil {
			res = append(res, s.getDifferenceProfit())
		}
		return res
	}
}

func main() {}
