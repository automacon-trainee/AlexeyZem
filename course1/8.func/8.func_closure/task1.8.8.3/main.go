package main

import (
	"fmt"
)

func adder(initial int) func(int) int {
	return func(x int) int {
		return x + initial
	}
}

func main() {
	initial := 2
	addTwo := adder(initial)
	num := 6
	fmt.Println(addTwo(num))
}
