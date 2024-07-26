package main

import (
	"fmt"
)

func CheckDiscount(price, discount float64) (float64, error) {
	if discount < 0 || discount > 50 {
		return 0.0, fmt.Errorf("too much discount")
	}
	newPrice := price * (1 - (discount / 100))
	return newPrice, nil
}

func main() {
	price := 100.0
	discount := 5.0
	fmt.Println(CheckDiscount(price, discount))
}
