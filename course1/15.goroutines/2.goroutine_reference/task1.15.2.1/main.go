// Это задача на исправление, поэтому линтер не проходит в некоторых местах.

package main

import (
	"fmt"
	"os"
	"time"
)

var ch = make(chan int)

const (
	one = iota + 1
	two
	three
	four
	five
)

func main() {
	numbers := []int{one, two, three, four, five}
	storeNumbers(numbers)
	print(ch)
}

func print(data chan int) {
	if len(os.Getenv("DEBUG")) != 0 {
		return
	}
	go func() {
		time.Sleep(1 * time.Second)
		close(ch)
	}()
	for v := range data {
		fmt.Println(v)
	}
}

func storeNumbers(data []int) {
	for _, n := range data {
		go func(n int) {
			go write(n)
		}(n)
	}
}

func write(n int) {
	ch <- n
}
