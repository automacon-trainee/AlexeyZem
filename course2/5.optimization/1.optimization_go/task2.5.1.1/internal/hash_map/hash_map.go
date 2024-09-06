package hash_map

import (
	"hash"
	"hash/crc32"
	"hash/crc64"

	"hashMap/internal"
)

// RehashCoefficientCapacity must be < RehashCoefficientLength
const (
	RehashCoefficientCapacity = 3
	RehashCoefficientLength   = 4
)

type HashMap struct {
	Container internal.Container
	hasher    hash.Hash
}

func (h *HashMap) Get(key string) (any, bool) {
	if RehashCoefficientCapacity*h.Container.Capacity() <= h.Container.Length()*RehashCoefficientLength {
		h.rehash()
	}
	return h.Container.Get(key, &h.hasher)
}

func (h *HashMap) Set(key string, val any) {
	if RehashCoefficientCapacity*h.Container.Capacity() <= h.Container.Length()*RehashCoefficientLength {
		h.rehash()
	}
	h.Container.Set(key, val, &h.hasher)
}

func (h *HashMap) rehash() {
	h.hasher.Reset()
	newCont := h.Container.Rehash(&h.hasher)
	h.Container = newCont
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

func NewHashMap(cont internal.Container, opt Option) *HashMap {
	res := &HashMap{Container: cont}
	opt(res)
	return res
}
