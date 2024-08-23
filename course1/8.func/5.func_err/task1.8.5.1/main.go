package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func CalculatePercentageChange(initalValue, finalValue string) (float64, error) {
	if initalValue == "0" {
		return 0.0, errors.New("inital value is zero")
	}
	finValue, err := strconv.ParseFloat(finalValue, 64)
	if err != nil {
		return 0, err
	}
	initValue, err := strconv.ParseFloat(initalValue, 64)
	if err != nil {
		return 0, err
	}

	fromNumToPercent := 100.0

	return math.Abs((finValue/initValue - 1) * fromNumToPercent), nil
}

func main() {
	initialValue := "10.0"
	finalValue := "1"
	res, err := CalculatePercentageChange(initialValue, finalValue)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println(res)
	}
}
