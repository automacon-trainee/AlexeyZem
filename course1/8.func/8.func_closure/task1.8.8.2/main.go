package main

import (
	"fmt"
)

func multiplier(factor float64) func(float64) float64 {
	return func(x float64) float64 {
		return x * factor
	}
}

func main() {
	factor := 2.5
	num := 10.0
	m := multiplier(factor)
	fmt.Println(m(num))
}
