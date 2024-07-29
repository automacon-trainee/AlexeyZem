package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	data []int
	want []int
}

func TestRemoveExtraMemory(t *testing.T) {
	first := make([]int, 3, 10)
	first[0] = 1
	first[1] = 2
	first[2] = 3
	tests := []TestCase{
		{first, []int{1, 2, 3}},
		{[]int{1, 2}, []int{1, 2}},
	}
	for _, test := range tests {
		res := RemoveExtraMemory(test.data)
		if !reflect.DeepEqual(res, test.want) {
			t.Errorf("RemoveExtraMemory(%v) = %v, want %v", test.data, res, test.want)
		}
	}
}
