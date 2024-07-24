package main

import "fmt"

func Dereference(a *int) int {
	if a != nil {
		return *a
	}
	return 0
}

func Sum(a, b *int) int {
	if a != nil && b != nil {
		return *a + *b
	}
	return 0
}

func main() {
	a := 5
	b := 10
	fmt.Println(Dereference(&a))
	fmt.Println(Sum(&a, &b))
}
