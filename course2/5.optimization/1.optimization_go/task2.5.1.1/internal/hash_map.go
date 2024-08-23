package internal

import (
	"encoding/binary"
	"hash"
	"hash/crc32"
	"hash/crc64"
)

// RehashCoefficientCapacity must be < RehashCoefficientLength
const (
	RehashCoefficientCapacity = 3
	RehashCoefficientLength   = 4
)

type Container interface {
	Set(key string, val any, hasher *hash.Hash)
	Get(key string, hasher *hash.Hash) (any, bool)
	rehash(hasher *hash.Hash) Container
	length() int
	capacity() int
}

type HashMap struct {
	Container Container
	hasher    hash.Hash
}

func (h *HashMap) Get(key string) (any, bool) {
	if RehashCoefficientCapacity*h.Container.capacity() <= h.Container.length()*RehashCoefficientLength {
		h.rehash()
	}
	return h.Container.Get(key, &h.hasher)
}

func (h *HashMap) Set(key string, val any) {
	if RehashCoefficientCapacity*h.Container.capacity() <= h.Container.length()*RehashCoefficientLength {
		h.rehash()
	}
	h.Container.Set(key, val, &h.hasher)
}

func (h *HashMap) rehash() {
	h.hasher.Reset()
	newCont := h.Container.rehash(&h.hasher)
	h.Container = newCont
}

func Hash(key string, hasher *hash.Hash) int {
	(*hasher).Reset()
	(*hasher).Write([]byte(key))
	return int(binary.BigEndian.Uint16((*hasher).Sum(nil)))
}

type Option func(h *HashMap)

func WithHashCRC64() Option {
	return func(h *HashMap) {
		h.hasher = crc64.New(crc64.MakeTable(crc64.ECMA))
	}
}

func WithHashCRC32() Option {
	return func(h *HashMap) {
		h.hasher = crc32.NewIEEE()
	}
}

func NewHashMap(cont Container, opt Option) *HashMap {
	res := &HashMap{Container: cont}
	opt(res)
	return res
}
