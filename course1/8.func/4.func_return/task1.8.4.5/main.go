package main

import (
	"fmt"
)

func CalculatePercentageChange(initalValue, finalValue float64) float64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Zero value")
		}
	}()
	fromNumToPercent := 100.0
	if finalValue > initalValue {
		return (finalValue/initalValue - 1) * fromNumToPercent
	}
	return (finalValue/initalValue - 1) * -fromNumToPercent
}

func main() {
	initialValue := 10.0
	finalValue := 10.6
	fmt.Println(CalculatePercentageChange(initialValue, finalValue))
}
