package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"

	"redis/entities"
)

const expiration = time.Minute * 30

type CacherClient interface {
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(key string) *redis.StringCmd
}

type CacherRedis struct {
	client CacherClient
}

func NewCacherRedis(client CacherClient) *CacherRedis {
	return &CacherRedis{client: client}
}

func (c *CacherRedis) Set(key string, value *entities.User) error {
	return c.client.Set(key, *value, expiration).Err()
}

func (c *CacherRedis) Get(key string) (entities.User, error) {
	val, err := c.client.Get(key).Result()
	if errors.Is(err, redis.Nil) {
		return entities.User{}, fmt.Errorf(`not found by key "%s"`, key)
	}
	var res entities.User
	err = json.Unmarshal([]byte(val), &res)
	return res, err
}
