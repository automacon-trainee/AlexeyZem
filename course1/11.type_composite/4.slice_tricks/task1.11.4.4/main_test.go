package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	data []int
	idx  int
	want []int
}

func TestRemoveIDX(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 2, 3}, 2, []int{1, 2}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3}, -1, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 0, []int{2, 3}},
	}
	for _, test := range tests {
		res := RemoveIDX(test.data, test.idx)
		if !reflect.DeepEqual(res, test.want) {
			t.Errorf("RemoveIDX(%v, %v) = %v, want %v", test.data, test.idx, res, test.want)
		}
	}
}
