package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	data  []int
	start int
	end   int
	want  []int
}

func TestCut(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 2, 3, 4, 5}, 1, 3, []int{2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, -1, 3, []int{}},
		{[]int{1, 2, 3, 4, 5}, 1, 10, []int{}},
		{[]int{1, 2, 3, 4, 5}, 4, 1, []int{}},
		{[]int{1, 2, 3, 4, 5}, 0, 4, []int{1, 2, 3, 4, 5}},
	}
	for _, test := range tests {
		res := Cut(test.data, test.start, test.end)
		if !reflect.DeepEqual(res, test.want) {
			t.Errorf("Cut(%v, %v, %v) = %v, want %v", test.data, test.start, test.end, res, test.want)
		}
	}
}
