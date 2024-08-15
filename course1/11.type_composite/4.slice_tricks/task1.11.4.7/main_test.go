package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	data   []int
	wantEl int
	wantSl []int
}

func TestPop(t *testing.T) {
	tests := []TestCase{
		{[]int{1, 2, 3}, 1, []int{2, 3}},
		{[]int{}, 0, []int{}},
	}
	for _, test := range tests {
		el, arr := Pop(test.data)
		if el != test.wantEl || !reflect.DeepEqual(arr, test.wantSl) {
			t.Errorf("Pop(%v) => %v, %v want %v %v", test.data, el, arr, test.wantEl, test.wantSl)
		}
	}
}
