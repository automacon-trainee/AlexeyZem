package main

import (
	"fmt"
)

func CreateCounter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	counter := CreateCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
