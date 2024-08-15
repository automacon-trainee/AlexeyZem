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

func TestAppendInt(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 2, 3}, []int{0, 0}, []int{1, 2, 3, 0, 0}},
		{[]int{}, []int{1}, []int{1}},
	}
	for _, test := range tests {
		result := appendInt(test.data, test.newData...)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("appendInt(%v, %v) => %v, want %v", test.data, test.newData, result, test.want)
		}
	}
}
