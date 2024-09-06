package main

import (
	"fmt"
)

func PrintNumbers(numbers ...int) {
	for _, n := range numbers {
		fmt.Println(n)
	}
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	PrintNumbers(sl...)
}
