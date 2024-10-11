package internal

import (
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type mockRepo map[string]string

const expiration = time.Minute * 10

type SomeRepositoryImpl struct {
	repo mockRepo // imitation repository
}

func NewSomeRepository() *SomeRepositoryImpl {
	return &SomeRepositoryImpl{repo: make(map[string]string)}
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

type CacherClient interface {
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(key string) *redis.StringCmd
}

type SomeRepository interface {
	GetData(key string) string
}

type SomeRepositoryProxyImpl struct {
	someRepo SomeRepository
	cache    CacherClient
}

func NewSomeRepositoryProxy(client CacherClient, origRepo SomeRepository) SomeRepository {
	return &SomeRepositoryProxyImpl{someRepo: origRepo, cache: client}
}

func (s *SomeRepositoryProxyImpl) GetData(key string) string {
	val, err := s.cache.Get(key).Result()
	if errors.Is(err, redis.Nil) {
		val = s.someRepo.GetData(key)
		s.cache.Set(key, val, expiration)
	}
	if err != nil {
		log.Println(err)
		return ""
	}
	return val
}
