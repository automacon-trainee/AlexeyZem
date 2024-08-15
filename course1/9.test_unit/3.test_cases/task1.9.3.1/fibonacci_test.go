package main

import (
	"testing"
)

type testCase struct {
	data   int
	answer int
}

func TestFibonacci(t *testing.T) {
	testCases := []testCase{
		{data: 5, answer: 5},
		{data: 6, answer: 8},
		{data: 7, answer: 13},
		{data: 8, answer: 21},
		{data: 9, answer: 34},
		{data: 10, answer: 55},
	}

	for _, testingCase := range testCases {
		res := Fibonacci(testingCase.data)
		if res != testingCase.answer {
			t.Errorf("Unexpected result: %d instead of %d", res, testingCase.answer)
		}
	}
}
