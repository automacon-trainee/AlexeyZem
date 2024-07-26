package main

import (
	"fmt"
)

func FindMaxAndMin(n ...int) (max, min int) {
	max, min = n[0], n[0]
	for _, v := range n {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}

func main() {
	max, min := FindMaxAndMin(1, 2, 3)
	fmt.Println(max, min)
}
