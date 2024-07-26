package main

import (
	"fmt"
)

func Factorial(n int) int {
	if n < 3 {
		return n
	}
	return n * Factorial(n-1)
}

func main() {
	num := 6
	fmt.Println(Factorial(num))
}
