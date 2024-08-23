package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// User Струкутура, описывающая пользователя
type User struct {
	ID   int
	Name string
}

// Cache потокобезопасный кэш, который позыоляет работать с обьектом типа User
type Cache struct {
	Data map[string]*User
	mu   sync.RWMutex
}

// NewCache - функция, которая создает новый кэш
func NewCache() *Cache {
	return &Cache{
		Data: make(map[string]*User),
		mu:   sync.RWMutex{},
	}
}

// Set метод кэша, записывает в кэш нового пользователя
func (c *Cache) Set(key string, user *User) {
	c.mu.Lock()
	c.Data[key] = user
	c.mu.Unlock()
}

// Get метод кэша, получает пользователя по указанному ключу
func (c *Cache) Get(key string) *User {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Data[key]
}

// функция, обьединяет строки через ":", формирую таким образом ключ кэша
func keyBuilder(keys ...string) string {
	return strings.Join(keys, ":")
}

// GetUser функция, преобразующая интерфейс в User, если это возможно
// linter need any
func GetUser(i any) *User {
	return i.(*User)
}

func main() {
	cache := NewCache()    // Create new Cache
	count := 10            // count of goroutines and users
	wg := sync.WaitGroup{} // for sync goroutines
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			cache.Set(keyBuilder("user", strconv.Itoa(i)), &User{ID: i, Name: fmt.Sprintf("user-%d", i)}) // set user in cache
			wg.Done()
		}(i)
	}
	wg.Wait() // wait until all users are in cache
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(i int) {
			raw := cache.Get(keyBuilder("user", strconv.Itoa(i))) // get user from cache with this key
			fmt.Println(GetUser(raw))                             // print user and transform from interface{} to *User
			wg.Done()
		}(i)
	}
	wg.Wait() // wait until all users are printed
}
