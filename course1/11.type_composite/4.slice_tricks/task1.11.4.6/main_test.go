package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	data    []int
	newData []int
	want    []int
}

func TestInsertToStart(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 2, 3}, []int{3, 2, 1}, []int{3, 2, 1, 1, 2, 3}},
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{[]int{}, []int{2}, []int{2}},
		{[]int{}, []int{}, []int{}},
	}
	for _, test := range tests {
		res := InsertToStart(test.data, test.newData...)
		if !reflect.DeepEqual(res, test.want) {
			t.Errorf("InsertToStart(%v,%v) = %v, want %v", test.data, test.newData, test.want, res)
		}
	}

}
