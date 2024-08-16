package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func concurrentSafeCounter() int {
	counter := Counter{}
	count := 10000
	for i := 0; i < count; i++ {
		go func() {
			counter.Incr()
		}()
	}
	time.Sleep(time.Second)
	return counter.value
}

func main() {
	fmt.Println(concurrentSafeCounter())
}
