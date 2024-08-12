package main

import (
	"fmt"
	"math"
)

func Floor(x float64) float64 {
	return math.Floor(x)
}

func main() {
	fmt.Println(Floor(2.9))
}
