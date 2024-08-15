package main

import (
	"reflect"
	"testing"
)

type TestCases struct {
	data  []int
	start int
	end   int
	want  []int
}

func TestGetSubSlice(t *testing.T) {
	tests := []TestCases{
		{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, start: 2, end: 6, want: []int{3, 4, 5, 6}},
		{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, start: 2, end: 10, want: []int{3, 4, 5, 6, 7, 8, 9, 10}},
	}
	for _, testCase := range tests {
		res := getSubSlice(testCase.data, testCase.start, testCase.end)
		if !reflect.DeepEqual(res, testCase.want) {
			t.Errorf("getSubSlice(%v, %v, %v): want %v, got %v", testCase.data, testCase.start, testCase.end, testCase.want, res)
		}
	}
}
