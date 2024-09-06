package slice

import (
	"hash"

	"hashMap/internal"
)

type HashMapSlice struct {
	hasher hash.Hash
	data   []*internal.Data
	length int
}

type Slice struct {
	data []*internal.Data
	len  int
	cap  int
}

func (s *Slice) Length() int {
	return s.len
}

func (s *Slice) Capacity() int {
	return s.cap
}

func (s *Slice) Rehash(hasher *hash.Hash) internal.Container {
	if s.Capacity() == 0 {
		return &Slice{make([]*internal.Data, 3), 0, 3}
	}
	data := make([]*internal.Data, 2*s.Capacity())
	for _, d := range s.data {
		if d != nil {
			hashVal := internal.Hash(d.Key, hasher)
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
	s.cap = 2 * s.Capacity()
	s.data = data
	return s
}

func (s *Slice) Set(key string, value any, hasher *hash.Hash) {
	hashVal := internal.Hash(key, hasher)
	hashVal %= s.Capacity()

	for hashVal < len(s.data) && s.data[hashVal] != nil {
		if s.data[hashVal].Key == key {
			s.data[hashVal].Value = value
			return
		}
		hashVal++
		if hashVal == len(s.data) {
			hashVal = 0
		}
	}
	s.data[hashVal] = &internal.Data{Key: key, Value: value}
	s.len++
}

func (s *Slice) Get(key string, hasher *hash.Hash) (any, bool) {
	if s.Length() == 0 {
		return nil, false
	}
	hashVal := internal.Hash(key, hasher)
	hashVal %= s.Capacity()
	for hashVal < len(s.data) && s.data[hashVal] != nil {
		if s.data[hashVal].Key == key {
			return s.data[hashVal].Value, true
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
		data: make([]*internal.Data, len),
		len:  0,
		cap:  len,
	}
}
