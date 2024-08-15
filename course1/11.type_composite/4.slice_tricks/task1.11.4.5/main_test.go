package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	data []int
	div  int
	want []int
}

func TestFilterDividers(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 2, 3}, 3, []int{3}},
		{[]int{1, 2, 3, 4}, 2, []int{2, 4}},
	}
	for _, test := range tests {
		res := FilterDividers(test.data, test.div)
		if !reflect.DeepEqual(res, test.want) {
			t.Errorf("FilterDividers(%v, %v) => %v, want %v", test.data, test.div, res, test.want)
		}
	}
}
