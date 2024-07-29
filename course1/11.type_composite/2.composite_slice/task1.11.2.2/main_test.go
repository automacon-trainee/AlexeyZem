package main

import (
	"testing"
)

type TestCase struct {
	data []int
	want int
}

func TestMaxDifference(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 2, 3}, 2},
		{[]int{1, 2, 3, 5, 0, 10, 3}, 10},
	}
	for _, test := range tests {
		res := MaxDifference(test.data)
		if res != test.want {
			t.Errorf("MaxDifference(%v) = %d, want %d", test.data, res, test.want)
		}
	}
}
