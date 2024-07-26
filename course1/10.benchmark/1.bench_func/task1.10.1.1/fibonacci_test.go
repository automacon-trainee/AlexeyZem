package main

import (
	"math/rand"
	"testing"
)

func BenchmarkFibonacci(b *testing.B) {
	data := generateData(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fibonacci(data[i])
	}
}

func generateData(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = rand.Intn(100)
	}
	return res
}

type testCase struct {
	data int
	answ int
}

func TestFibonacci(t *testing.T) {
	tests := []testCase{
		{data: 0, answ: 0},
		{data: 1, answ: 1},
		{data: 2, answ: 1},
		{data: 3, answ: 2},
		{data: 4, answ: 3},
		{data: 5, answ: 5},
		{data: 6, answ: 8},
		{data: 7, answ: 13},
		{data: 8, answ: 21},
		{data: 9, answ: 34},
		{data: 10, answ: 55},
	}
	for _, test := range tests {
		if result := Fibonacci(test.data); result != test.answ {
			t.Errorf("Fibonacci(%d) = %d, want %d", test.data, result, test.answ)
		}
	}
}
