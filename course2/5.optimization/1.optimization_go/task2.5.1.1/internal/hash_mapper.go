package internal

import (
	"hash"
)

type Data struct {
	Key   string
	Value any
}

type HashMapper interface {
	Set(key string, value any)
	Get(key string) (any, bool)
}

type Container interface {
	Set(key string, val any, hasher *hash.Hash)
	Get(key string, hasher *hash.Hash) (any, bool)
	Rehash(hasher *hash.Hash) Container
	Length() int
	Capacity() int
}
