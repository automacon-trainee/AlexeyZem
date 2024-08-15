package internal

import (
	"strconv"
	"testing"
)

func BenchmarkHashMapSliceCRC64(b *testing.B) {
	hashSl := NewHashMapSlice(1, WithHashCRC64())
	b.ResetTimer()
	for i := 0; i < 50000; i++ {
		hashSl.Set(strconv.Itoa(i), strconv.Itoa(i))
		hashSl.Get(strconv.Itoa(i))
	}
}

func BenchmarkHashMapSliceCRC32(b *testing.B) {
	hashSl := NewHashMapSlice(1, WithHashCRC32())
	b.ResetTimer()
	for i := 0; i < 50000; i++ {
		hashSl.Set(strconv.Itoa(i), strconv.Itoa(i))
		hashSl.Get(strconv.Itoa(i))
	}
}

func BenchmarkHashMapList(b *testing.B) {
	hashL := NewHashMapList()
	b.ResetTimer()
	for i := 0; i < 50000; i++ {
		hashL.Set(strconv.Itoa(i), strconv.Itoa(i))
		hashL.Get(strconv.Itoa(i))
	}
}
