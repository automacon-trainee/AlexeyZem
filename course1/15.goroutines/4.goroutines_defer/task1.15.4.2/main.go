package main

import (
	"fmt"
)

func main() {
	a, b := 8, 13
	fmt.Println(*(TestDefer(&a, &b)))
}
func TestDefer(a, b *int) *int {
	var c int
	defer func() {
		c = multiply(*a, *b)
	}()
	c = sum(*a, *b)
	return &c
}
func sum(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}
