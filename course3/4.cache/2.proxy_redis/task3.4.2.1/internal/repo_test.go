package internal

import (
	"testing"

	"github.com/go-redis/redis"
)

func TestSomeRepositoryImpl(t *testing.T) {
	repoImpl := NewSomeRepository()
	val := repoImpl.GetData("some")
	if val != "" {
		t.Errorf("SomeRepositoryImpl.GetData() should return empty string")
	}
	repoImpl.SetData("some", "value")
	val = repoImpl.GetData("some")
	if val != "value" {
		t.Errorf("SomeRepositoryImpl.GetData() should return 'value'")
	}
}

func TestSomeRepositoryProxyImpl(t *testing.T) {
	repoImpl := NewSomeRepository()
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	proxyRepo := NewSomeRepositoryProxy(client, repoImpl)
	val := proxyRepo.GetData("some")
	if val != "" {
		t.Errorf("SomeRepositoryProxyImpl.GetData() should return empty string")
	}
	repoImpl.SetData("some2", "value")
	val = proxyRepo.GetData("some2")
	if val != "value" {
		t.Errorf("SomeRepositoryProxyImpl.GetData() should return 'value'")
	}
	val = proxyRepo.GetData("some2")
	if val != "value" {
		t.Errorf("SomeRepositoryProxyImpl.GetData() should return 'value' with redis too")
	}
}
