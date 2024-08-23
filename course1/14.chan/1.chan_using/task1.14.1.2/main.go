package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func generateData(n int) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			num, err := rand.Int(rand.Reader, big.NewInt(big.MaxExp))
			if err != nil {
				panic(err)
			}
			out <- int(num.Int64())
		}
		close(out)
	}()
	return out
}

func main() {
	count := 90
	res := generateData(count)
	for num := range res {
		fmt.Println(num)
	}
}
