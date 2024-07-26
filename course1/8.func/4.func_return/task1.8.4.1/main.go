package main

import (
	"fmt"
)

func DevideAndRemainder(a, b int) (dev, rem int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("check zero argument")
		}
	}()
	dev = a / b
	rem = a % b
	return dev, rem
}

func main() {
	a := 0
	b := 20
	dev, rem := DevideAndRemainder(b, a)
	fmt.Printf("Частное: %d, Остаток: %d\n", dev, rem)
}
