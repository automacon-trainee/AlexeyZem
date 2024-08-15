package main

import (
	"encoding/binary"
	"hash"
	"hash/crc32"
	"hash/crc64"
	"time"

	"github.com/howeyc/crc16"
)

type HashMapper interface {
	Set(key string, value any)
	Get(key string) (any, bool)
}

func MeasureTime(task func()) time.Duration {
	start := time.Now()
	task()
	return time.Since(start)
}

type Data struct {
	key string
	val any
}

type HashMap struct {
	hasher hash.Hash
	data   []*Data
	length int
}

func (h *HashMap) rehash() {
	h.hasher.Reset()
	if len(h.data) == 0 {
		h.data = make([]*Data, 3)
	}
	data := make([]*Data, 2*len(h.data))
	for _, d := range h.data {
		if d != nil {
			hashVal := Hash(h, d.key)
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

func (h *HashMap) Set(key string, value any) {
	hashVal := Hash(h, key)
	if 3*cap(h.data) <= 4*h.length || cap(h.data) == 0 {
		h.rehash()
	}
	hashVal %= len(h.data)

	for hashVal < len(h.data) && h.data[hashVal] != nil {
		if h.data[hashVal].key == key {
			h.data[hashVal].val = value
			return
		}
		hashVal++
		if hashVal == len(h.data) {
			hashVal = 0
		}
	}
	h.data[hashVal] = &Data{key: key, val: value}
	h.length++
}

func (h *HashMap) Get(key string) (any, bool) {
	if h.length == 0 {
		return nil, false
	}
	hashVal := Hash(h, key)
	hashVal %= len(h.data)
	for hashVal < len(h.data) && h.data[hashVal] != nil {
		if h.data[hashVal].key == key {
			return h.data[hashVal].val, true
		}
		hashVal++
		if hashVal == len(h.data) {
			hashVal = 0
		}
	}
	return nil, false
}

func Hash(h *HashMap, key string) int {
	h.hasher.Reset()
	h.hasher.Write([]byte(key))
	return int(binary.BigEndian.Uint16(h.hasher.Sum(nil)))
}

type Option func(*HashMap)

func NewHashMap(size int, opt Option) *HashMap {
	data := make([]*Data, size)
	h := &HashMap{data: data, length: 0}
	opt(h)
	return h
}

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

func WithHashCRC16() Option {
	return func(h *HashMap) {
		h.hasher = crc16.New(crc16.MakeTable(crc16.SCSI))
	}
}

func main() {}
