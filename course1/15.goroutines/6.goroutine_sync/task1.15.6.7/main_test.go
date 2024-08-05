package main

import (
	"sync"
	"testing"
)

type Person struct {
	Age int
}

var personPool = sync.Pool{
	New: func() interface{} {
		return &Person{}
	},
}

func BenchmarkWithPool(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := personPool.Get().(*Person)
		personPool.Put(p)
	}
}

func BenchmarkWithoutPool(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := Person{Age: i}
		_ = p
	}
}
