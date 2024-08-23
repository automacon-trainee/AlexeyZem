package main

import (
	"fmt"
	"sync"
)

func mergeChan(mergeTo chan int, from ...chan int) {
	wg := sync.WaitGroup{}
	for _, ch := range from {
		wg.Add(1)
		go func(ch chan int) {
			defer wg.Done()
			for num := range ch {
				mergeTo <- num
			}
		}(ch)
	}
	wg.Wait()
	close(mergeTo)
}

func mergeChan2(chans ...chan int) chan int {
	res := make(chan int)
	go func() {
		wg := sync.WaitGroup{}
		for _, ch := range chans {
			wg.Add(1)
			go func(ch chan int) {
				defer wg.Done()
				for num := range ch {
					res <- num
				}
			}(ch)
		}
		wg.Wait()
		close(res)
	}()
	return res
}

func generateChan(n int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func main() {
	res := make(chan int)
	count := 10
	someChan := make([]chan int, count)
	for i := 0; i < count; i++ {
		someChan[i] = generateChan(count)
	}

	go mergeChan(res, someChan...)
	fmt.Println("Start first")
	for num := range res {
		fmt.Println(num)
	}
	fmt.Println()
	fmt.Println("Start second")
	for i := 0; i < count; i++ {
		someChan[i] = generateChan(count)
	}
	ch := mergeChan2(someChan...)
	for num := range ch {
		fmt.Println(num)
	}
}
