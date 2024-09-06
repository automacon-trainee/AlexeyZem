package main

import (
	"fmt"
)

func DevideAndRemainder(a, b int) (dev, rem int) {
	if b == 0 {
		return 0, 0
	}
	dev = a / b
	rem = a % b
	return
}

func main() {
	a := 0
	b := 20
	dev, rem := DevideAndRemainder(b, a)
	fmt.Printf("Частное: %d, Остаток: %d\n", dev, rem)
}
