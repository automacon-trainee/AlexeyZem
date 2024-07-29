package main

import (
	"testing"
)

type TestCase struct {
	str  string
	want int
}

func TestCountVowels(t *testing.T) {
	testCases := []TestCase{
		{"Привет, мир", 3},
		{"Hello, world!", 3},
		{"Привет, world, it is база", 7},
	}
	for _, testCase := range testCases {
		res := CountVowels(testCase.str)
		if res != testCase.want {
			t.Errorf("countVowels(%s)=%d, want %d", testCase.str, res, testCase.want)
		}
	}
}
