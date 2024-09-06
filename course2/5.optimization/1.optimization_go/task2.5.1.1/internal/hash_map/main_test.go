package hash_map

import (
	"strconv"
	"testing"

	"hashMap/internal/list"
	"hashMap/internal/slice"
)

func BenchmarkHashMapSliceCRC64(b *testing.B) {
	hashSl := NewHashMap(slice.NewSlice(1), WithHashCRC64())
	b.ResetTimer()
	for i := 0; i < 10000; i++ {
		hashSl.Set(strconv.Itoa(i), strconv.Itoa(i))
		hashSl.Get(strconv.Itoa(i))
	}
}

func BenchmarkHashMapSliceCRC32(b *testing.B) {
	hashSl := NewHashMap(slice.NewSlice(1), WithHashCRC32())
	b.ResetTimer()
	for i := 0; i < 10000; i++ {
		hashSl.Set(strconv.Itoa(i), strconv.Itoa(i))
		hashSl.Get(strconv.Itoa(i))
	}
}

func BenchmarkHashMapListCRC64(b *testing.B) {
	hashL := NewHashMap(list.CreateNewList(1), WithHashCRC64())
	b.ResetTimer()
	for i := 0; i < 10000; i++ {
		hashL.Set(strconv.Itoa(i), strconv.Itoa(i))
		hashL.Get(strconv.Itoa(i))
	}
}
func BenchmarkHashMapListCRC32(b *testing.B) {
	hashL := NewHashMap(list.CreateNewList(1), WithHashCRC32())
	b.ResetTimer()
	for i := 0; i < 10000; i++ {
		hashL.Set(strconv.Itoa(i), strconv.Itoa(i))
		hashL.Get(strconv.Itoa(i))
	}
}
