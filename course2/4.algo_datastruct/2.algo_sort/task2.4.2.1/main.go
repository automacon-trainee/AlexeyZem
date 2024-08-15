package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit"
)

type Product struct {
	Name      string
	Price     float64
	CreatedAt time.Time
	Count     int
}

func (p *Product) String() string {
	return fmt.Sprintf("Name: %s, Price: %.2f, Count: %v", p.Name, p.Price, p.Count)
}

func generateProducts(num int) []Product {
	gofakeit.Seed(time.Now().UnixNano())
	products := make([]Product, num)
	for i := range products {
		products[i] = Product{
			Name:      gofakeit.Name(),
			Price:     gofakeit.Price(1.0, 100.0),
			CreatedAt: gofakeit.Date(),
			Count:     gofakeit.Number(1, 100),
		}
	}
	return products
}

type Sorting interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type ProductSorter struct {
	products []Product
	less     func(i, j int) bool
}

func (s *ProductSorter) Len() int {
	return len(s.products)
}
func (s *ProductSorter) Less(i, j int) bool {
	return s.less(i, j)
}
func (s *ProductSorter) Swap(i, j int) {
	s.products[i], s.products[j] = s.products[j], s.products[i]
}

func ByCreatedAt(products []Product) Sorting {
	return &ProductSorter{
		products: products,
		less: func(i, j int) bool {
			return products[i].CreatedAt.Unix() < products[j].CreatedAt.Unix()
		},
	}
}

func ByPrice(products []Product) Sorting {
	return &ProductSorter{
		products: products,
		less: func(i, j int) bool {
			return products[i].Price < products[j].Price
		},
	}
}

func ByCount(products []Product) Sorting {
	return &ProductSorter{
		products: products,
		less: func(i, j int) bool {
			return products[i].Count < products[j].Count
		},
	}
}

func main() {}
