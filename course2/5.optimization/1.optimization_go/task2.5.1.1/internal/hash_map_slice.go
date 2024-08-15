package internal

import (
	"encoding/binary"
	"hash"
	"hash/crc32"
	"hash/crc64"
)

type HashMapSlice struct {
	hasher hash.Hash
	data   []*Data
	length int
}

func (h *HashMapSlice) Hash(key string) int {
	h.hasher.Reset()
	h.hasher.Write([]byte(key))
	return int(binary.BigEndian.Uint16(h.hasher.Sum(nil)))
}

func (h *HashMapSlice) rehash() {
	h.hasher.Reset()
	if len(h.data) == 0 {
		h.data = make([]*Data, 3)
	}
	data := make([]*Data, 2*len(h.data))
	for _, d := range h.data {
		if d != nil {
			hashVal := h.Hash(d.key)
			hashVal %= len(data)
			for hashVal < len(data) && data[hashVal] != nil {
				hashVal++
				if hashVal == len(data) {
					hashVal = 0
				}
			}
			data[hashVal] = d
		}
	}
	h.data = data
}

func (h *HashMapSlice) Set(key string, value any) {
	hashVal := h.Hash(key)
	if 3*cap(h.data) <= 4*h.length || cap(h.data) == 0 {
		h.rehash()
	}
	hashVal %= len(h.data)

	for hashVal < len(h.data) && h.data[hashVal] != nil {
		if h.data[hashVal].key == key {
			h.data[hashVal].value = value
			return
		}
		hashVal++
		if hashVal == len(h.data) {
			hashVal = 0
		}
	}
	h.data[hashVal] = &Data{key: key, value: value}
	h.length++
}

func (h *HashMapSlice) Get(key string) (any, bool) {
	if h.length == 0 {
		return nil, false
	}
	if 3*cap(h.data) <= 4*h.length || cap(h.data) == 0 {
		h.rehash()
	}
	hashVal := h.Hash(key)
	hashVal %= len(h.data)
	for hashVal < len(h.data) && h.data[hashVal] != nil {
		if h.data[hashVal].key == key {
			return h.data[hashVal].value, true
		}
		hashVal++
		if hashVal == len(h.data) {
			hashVal = 0
		}
	}
	return nil, false
}

type Option func(slice *HashMapSlice)

func NewHashMapSlice(size int, opt Option) *HashMapSlice {
	data := make([]*Data, size)
	h := &HashMapSlice{data: data, length: 0}
	opt(h)
	return h
}

func WithHashCRC64() Option {
	return func(h *HashMapSlice) {
		h.hasher = crc64.New(crc64.MakeTable(crc64.ECMA))
	}
}

func WithHashCRC32() Option {
	return func(h *HashMapSlice) {
		h.hasher = crc32.NewIEEE()
	}
}
