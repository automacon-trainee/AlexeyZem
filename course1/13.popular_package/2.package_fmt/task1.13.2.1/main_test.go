package main

import (
	"testing"
)

type TestCase struct {
	nums     []int
	operator string
	want     string
}

func TestGenerateMathString(t *testing.T) {
	testCases := []TestCase{
		{nums: []int{1, 2, 3}, operator: "+", want: "1 + 2 + 3 = 6"},
		{nums: []int{1, 2, 3}, operator: "-", want: "1 - 2 - 3 = -4"},
		{nums: []int{1, 2, 3}, operator: "*", want: "1 * 2 * 3 = 6"},
		{nums: []int{1, 2, 3}, operator: "/", want: "1 / 2 / 3 = 0"},
		{nums: []int{1, 2, 3}, operator: "a", want: ""},
		{nums: []int{}, operator: "*", want: ""},
	}
	for _, tc := range testCases {
		res := generateMathStrings(tc.nums, tc.operator)
		if res != tc.want {
			t.Errorf("generateMathStrings(%v, %v) => %v, want %v", tc.nums, tc.operator, res, tc.want)
		}
	}
}
