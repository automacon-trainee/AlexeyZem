package main

import (
	"fmt"
	"math"
)

func Sin(x float64) float64 {
	return math.Sin(x)
}

func Cos(x float64) float64 {
	return math.Cos(x)
}

func main() {
	fmt.Println(Sin(2))
	fmt.Println(Cos(2))
}
