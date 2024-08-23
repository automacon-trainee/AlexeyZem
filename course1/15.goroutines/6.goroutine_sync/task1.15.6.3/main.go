package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	count int64
}

func (c *Counter) Incr() {
	atomic.AddInt64(&c.count, 1)
}

func (c *Counter) GetCount() int64 { return c.count }

func main() {
	c := new(Counter)
	count := 10000
	wg := sync.WaitGroup{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			c.Incr()
		}()
	}
	wg.Wait()
	fmt.Println(c.GetCount())
}
