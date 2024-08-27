package internal

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

const expiration = time.Minute * 30

type Cacher interface {
	Set(key string, value any) error
	Get(key string) (any, error)
}

type CacherRedis struct {
	client *redis.Client
}

func NewCacherRedis(client *redis.Client) *CacherRedis {
	return &CacherRedis{client: client}
}

func (c *CacherRedis) Set(key string, value any) error {
	return c.client.Set(key, value, expiration).Err()
}

func (c *CacherRedis) Get(key string) (any, error) {
	val, err := c.client.Get(key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf(`not found by key "%s"`, key)
	}
	return val, err
}
