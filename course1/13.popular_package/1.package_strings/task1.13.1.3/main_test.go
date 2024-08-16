package main

import (
	"testing"
)

type TestCase struct {
	length int
	want   int
}

func TestGenerateRandomString(t *testing.T) {
	testCases := []TestCase{
		{length: 10, want: 10},
		{length: 100, want: 100},
	}
	for _, testCase := range testCases {
		result := GenerateRandomString(testCase.length)
		if len([]rune(result)) != testCase.want {
			t.Errorf("GenerateRandomString(%v) = %v; want %v", testCase.length, result, testCase.want)
		}
	}
}
