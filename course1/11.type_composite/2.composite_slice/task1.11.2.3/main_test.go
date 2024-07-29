package main

import (
	"testing"
)

type TestCase struct {
	data []int
	want int
}

func TestFindSingleNumber(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 2, 1, 3, 2}, 3},
		{[]int{1, 1, 2}, 2},
		{[]int{1, 1}, 0},
	}
	for _, test := range tests {
		if got := findSingleNumber(test.data); got != test.want {
			t.Errorf("findSingleNumber(%v, %v): got %v, want %v", test.data, test.want, got, test.want)
		}
	}
}
