package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	m1   map[string]int
	m2   map[string]int
	want map[string]int
}

func TestMergeMaps(t *testing.T) {
	tests := []TestCase{
		{map[string]int{"a": 1}, map[string]int{"b": 2}, map[string]int{"a": 1, "b": 2}},
		{map[string]int{"a": 2}, map[string]int{"a": 2}, map[string]int{"a": 4}},
	}
	for _, test := range tests {
		res := mergeMaps(test.m1, test.m2)
		if !reflect.DeepEqual(res, test.want) {
			t.Errorf("mergeMaps(%v, %v) = %v, want %v", test.m1, test.m2, res, test.want)
		}
	}
}
