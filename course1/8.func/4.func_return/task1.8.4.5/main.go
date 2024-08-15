package main

import (
	"fmt"
	"math"
)

func CalculatePercentageChange(initalValue, finalValue float64) float64 {
	if initalValue == 0 {
		return 0
	}
	fromNumToPercent := 100.0
	return math.Abs((finalValue/initalValue - 1) * fromNumToPercent)
}

func main() {
	initialValue := 10.0
	finalValue := 16.6
	fmt.Println(CalculatePercentageChange(initialValue, finalValue))
}
