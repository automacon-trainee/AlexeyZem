package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	data    []int
	idx     int
	newData []int
	want    []int
}

func TestInsertAfterIndex(t *testing.T) {
	testCases := []TestCase{
		{[]int{1, 2, 3}, 1, []int{5, 7, 0}, []int{1, 2, 5, 7, 0, 3}},
		{[]int{1, 2}, 10, []int{5, 7, 0}, []int{}},
	}
	for _, testCase := range testCases {
		res := insertAfterIndex(testCase.data, testCase.idx, testCase.newData...)
		if !reflect.DeepEqual(res, testCase.want) {
			t.Errorf("insertAfterIndex(%v, %v) = %v, want %v", testCase.data, testCase.idx, res, testCase.want)
		}
	}
}
