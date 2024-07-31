package main

import (
	"testing"
)

type TestCase struct {
	data interface{}
	want string
}

func TestGetVariableType(t *testing.T) {
	testCases := []TestCase{
		{1, "int"},
		{1.1, "float64"},
		{true, "bool"},
		{"string", "string"},
		{[]int{}, "[]int"},
	}
	for _, testCase := range testCases {
		if got := getVariableType(testCase.data); got != testCase.want {
			t.Errorf("getVariableType(%v) = %v, want %v", testCase.data, got, testCase.want)
		}
	}
}
