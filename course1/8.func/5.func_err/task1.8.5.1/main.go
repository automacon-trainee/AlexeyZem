package main

import (
	"fmt"
	"strconv"
)

func CalculatePercentageChange(initalValue, finalValue string) (float64, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Zero value")
		}
	}()
	finValue, err := strconv.ParseFloat(finalValue, 64)
	if err != nil {
		return 0, err
	}
	initValue, err := strconv.ParseFloat(initalValue, 64)
	if err != nil {
		return 0, err
	}
	fromNumToPercent := 100.0
	if finalValue > initalValue {
		return (finValue/initValue - 1) * fromNumToPercent, nil
	}
	return (finValue/initValue - 1) * -fromNumToPercent, nil
}

func main() {
	initialValue := "10.0"
	finalValue := "25"
	res, err := CalculatePercentageChange(initialValue, finalValue)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println(res)
	}
}
