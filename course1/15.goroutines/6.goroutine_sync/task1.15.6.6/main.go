package main

import (
	"fmt"
	"strconv"
	"sync"
)

type Cache struct {
	Data sync.Map
}

func (c *Cache) Set(key string, value any) {
	c.Data.Store(key, value)
}

func (c *Cache) Get(key string) (any, bool) {
	return c.Data.Load(key)
}

func main() {
	cache := &Cache{}
	count := 10
	wg := sync.WaitGroup{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			defer wg.Done()
			cache.Set(strconv.Itoa(i), i)
		}(i)
	}
	wg.Wait()
	for i := 0; i < count+1; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(cache.Get(strconv.Itoa(i)))
		}(i)
	}
	wg.Wait()
}
