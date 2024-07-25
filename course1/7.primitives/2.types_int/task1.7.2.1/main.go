package main

import (
	"fmt"
	"log"
)

func calculate(a, b int) (sum, difference, product, quotient, remainder int) {
	sum = a + b
	difference = a - b
	product = a * b
	quotient = a / b
	remainder = a % b
	return sum, difference, product, quotient, remainder
}

func main() {
	var a, b, sum, difference, product, quotient, remainder int
	_, err := fmt.Scanln(&a, &b)
	if err != nil {
		log.Println("Wrong data:", err)
		return
	}

	sum, difference, product, quotient, remainder = calculate(a, b)
	fmt.Printf("a = %d, b = %d\nsum = %d, difference = %d, product = %d, quotient = %d, remainder = %d\n",
		a, b, sum, difference, product, quotient, remainder)
}
