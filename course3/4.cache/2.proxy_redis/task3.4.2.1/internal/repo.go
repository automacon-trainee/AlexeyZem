package internal

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type mockRepo map[string]string

const expiration = time.Minute * 10

type SomeRepository interface {
	GetData(key string) string
}

type SomeRepositoryImpl struct {
	repo mockRepo // imitation repository
}

func (s *SomeRepositoryImpl) GetData(key string) string {
	val, ok := s.repo[key]
	if ok {
		return val
	}
	return ""
}

func (s *SomeRepositoryImpl) SetData(key, val string) {
	s.repo[key] = val
}

func NewSomeRepository() *SomeRepositoryImpl {
	return &SomeRepositoryImpl{repo: make(map[string]string)}
}

type SomeRepositoryProxyImpl struct {
	someRepo SomeRepository
	cache    *redis.Client
}

func (s *SomeRepositoryProxyImpl) GetData(key string) string {
	val, err := s.cache.Get(key).Result()
	if errors.Is(err, redis.Nil) {
		val = s.someRepo.GetData(key)
		fmt.Println("hello from orig repo")
		s.cache.Set(key, val, expiration)
	}
	return val
}

func NewSomeRepositoryProxy(client *redis.Client, origRepo SomeRepository) SomeRepository {
	return &SomeRepositoryProxyImpl{someRepo: origRepo, cache: client}
}
