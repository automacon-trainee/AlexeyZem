package main

import (
	"fmt"
	"time"
)

func NotifyEvery(ticker *time.Ticker, timeout time.Duration, message string) chan string {
	c := make(chan string)
	stop := time.After(timeout)
	go func() {
	loop:
		for {
			select {
			case <-stop:
				break loop
			case <-ticker.C:
				c <- message
			}
		}
		close(c)
	}()

	return c
}

func main() {
	ticker := time.NewTicker(time.Second)

	data := NotifyEvery(ticker, 3*time.Second, "Таймер сработал")
	for v := range data {
		fmt.Println(v)
	}
	fmt.Println("Программа завершила работу")
}
