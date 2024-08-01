package main

import (
	"fmt"
	"time"
)

func timeout(timeout time.Duration) func() bool {
	return func() bool {
		ch := make(chan bool)
		defer close(ch)
		go func() {
			time.Sleep(time.Millisecond * 3) //some work
			ch <- true
		}()
		select {
		case <-ch:
			return true
		case <-time.After(timeout):
			return false
		}
	}
}

func main() {
	myFunc := timeout(time.Millisecond * 2)
	fmt.Println(myFunc())
}
