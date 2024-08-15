package main

import (
	"fmt"
	"math"
)

func CompareRoundedValues(a, b float64, decimalPlaces int) (isEqual bool, difference float64) {
	dec := float64(decimalPlaces) * 10
	a = math.Round(a*dec) / dec
	b = math.Round(b*dec) / dec
	isEqual = a == b
	difference = math.Abs(a - b)
	return
}

func main() {
	fmt.Println(CompareRoundedValues(2.757, 2.756, 2))
}
