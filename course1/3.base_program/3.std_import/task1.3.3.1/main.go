package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	result := math.Sqrt(x)
	return result
}

func main() {
	fmt.Println(sqrt(2))
}
