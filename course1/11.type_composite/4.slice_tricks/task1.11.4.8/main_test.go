package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	data      []int
	wantSlice []int
	wantFirst int
}

func TestShift(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 2, 3}, []int{3, 1, 2}, 3},
		{[]int{}, []int{}, 0},
		{[]int{1, 2, 3, 4, 5}, []int{5, 1, 2, 3, 4}, 5},
	}
	for _, test := range tests {
		first, res := Shift(test.data)
		if !reflect.DeepEqual(res, test.wantSlice) || first != test.wantFirst {
			t.Errorf("shift(%v) = %v, %v, want %v, %v", test.data, first, res, test.wantFirst, test.wantSlice)
		}
	}
}
