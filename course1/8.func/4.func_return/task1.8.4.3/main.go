package main

import (
	"fmt"
)

func CalculateStockValue(price float64, quantity int) (sum, one float64) {
	sum = price * float64(quantity)
	one = price
	return sum, one
}

func main() {
	price := 10.8
	quantity := 5
	fmt.Println(CalculateStockValue(price, quantity))
}
