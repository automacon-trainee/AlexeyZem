package main

import (
	"fmt"
	"math"
)

func hypotenuse(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}

func main() {
	var a, b = 3.0, 4.0
	fmt.Println(hypotenuse(a, b))
}
