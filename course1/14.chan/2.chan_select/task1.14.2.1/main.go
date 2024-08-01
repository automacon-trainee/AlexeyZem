package main

import (
	"fmt"
	"time"
)

func trySend(ch chan int, v int) bool {
	select {
	case ch <- v:
		return true
	default:
		return false
	}
}

func main() {
	ch := make(chan int, 1)
	go func() {
		ch <- 1
	}()
	time.Sleep(time.Millisecond)
	num := 10
	res := trySend(ch, num)
	fmt.Println(<-ch)
	fmt.Println(res)
}
