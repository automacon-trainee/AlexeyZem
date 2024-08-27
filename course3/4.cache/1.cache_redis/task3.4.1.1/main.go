// Реализуй  функционал,  который  позволит  выполнять  основные  операции  с  Redis,  такие  как
// установка  и
// получение значений по ключу.
// Описание задачи
// Задача состоит в создании кэширующего компонента Cacher на основе Redis

package main

import (
	"fmt"

	"github.com/go-redis/redis"

	"redis/entities"
	"redis/internal"
)

func main() {
	client := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
	cache := internal.NewCacherRedis(client)
	err := cache.Set("first", entities.User{ID: 1, Name: "John", Age: 20})
	if err != nil {
		panic(err)
	}
	val, err := cache.Get("first")
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
