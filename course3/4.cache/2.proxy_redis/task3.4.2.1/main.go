// Оптимизируй скорость запросов данных из внешнего источника путем кэширования.
// Задача — реализовать паттерн прокси для кэширования данных

package main

import (
	"fmt"

	"github.com/go-redis/redis"

	"proxy/internal"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	someRepo := internal.NewSomeRepository()
	someRepoProxy := internal.NewSomeRepositoryProxy(client, someRepo)
	someRepo.SetData("key", "value")
	val := someRepoProxy.GetData("key")
	fmt.Println("val:", val)
	val = someRepoProxy.GetData("key")
	fmt.Println("val:", val)
}
