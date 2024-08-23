package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type User struct {
	ID   int
	Name string
}

type Cache struct {
	Data map[string]*User
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		Data: make(map[string]*User),
		mu:   sync.RWMutex{},
	}
}

func (c *Cache) Set(key string, user *User) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Data[key] = user
}

func (c *Cache) Get(key string) *User {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Data[key]
}

func keyBuilder(keys ...string) string {
	return strings.Join(keys, ":")
}

func main() {
	cache := NewCache()
	count := 10
	wg := sync.WaitGroup{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			cache.Set(keyBuilder("user", strconv.Itoa(i)), &User{ID: i, Name: fmt.Sprintf("user-%d", i)})
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(*cache.Get(keyBuilder("user", strconv.Itoa(i))))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
