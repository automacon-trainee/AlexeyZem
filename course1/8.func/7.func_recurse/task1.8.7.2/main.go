package main

import (
	"fmt"
)

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	num := 6
	fmt.Println(Fibonacci(num))
}
