package internal

import (
	"hash"
)

type HashMapSlice struct {
	hasher hash.Hash
	data   []*Data
	length int
}

type Slice struct {
	data []*Data
	len  int
	cap  int
}

func (s *Slice) length() int {
	return s.len
}

func (s *Slice) capacity() int {
	return s.cap
}

func (s *Slice) rehash(hasher *hash.Hash) Container {
	if s.capacity() == 0 {
		return &Slice{make([]*Data, 3), 0, 3}
	}
	data := make([]*Data, 2*s.capacity())
	for _, d := range s.data {
		if d != nil {
			hashVal := Hash(d.key, hasher)
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
	s.cap = 2 * s.capacity()
	s.data = data
	return s
}

func (s *Slice) Set(key string, value any, hasher *hash.Hash) {
	hashVal := Hash(key, hasher)
	hashVal %= s.capacity()

	for hashVal < len(s.data) && s.data[hashVal] != nil {
		if s.data[hashVal].key == key {
			s.data[hashVal].value = value
			return
		}
		hashVal++
		if hashVal == len(s.data) {
			hashVal = 0
		}
	}
	s.data[hashVal] = &Data{key: key, value: value}
	s.len++
}

func (s *Slice) Get(key string, hasher *hash.Hash) (any, bool) {
	if s.length() == 0 {
		return nil, false
	}
	hashVal := Hash(key, hasher)
	hashVal %= s.capacity()
	for hashVal < len(s.data) && s.data[hashVal] != nil {
		if s.data[hashVal].key == key {
			return s.data[hashVal].value, true
		}
		hashVal++
		if hashVal == len(s.data) {
			hashVal = 0
		}
	}
	return nil, false
}

func NewSlice(len int) *Slice {
	return &Slice{
		data: make([]*Data, len),
		len:  0,
		cap:  len,
	}
}
