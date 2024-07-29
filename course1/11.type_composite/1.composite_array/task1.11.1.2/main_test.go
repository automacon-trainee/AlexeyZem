package main

import (
	"testing"
)

type TestCases struct {
	data   [8]int
	answer [8]int
}

func TestSort(t *testing.T) {
	tests := []TestCases{
		{data: [8]int{1, 2, 3, 4, 5, 8, 7, 6}, answer: [8]int{1, 2, 3, 4, 5, 6, 7, 8}},
		{data: [8]int{1, 2, 3, 4, 5, 6, 7, 8}, answer: [8]int{1, 2, 3, 4, 5, 6, 7, 8}},
	}
	for _, testCase := range tests {
		res := sortAsc(testCase.data)
		if res != testCase.answer {
			t.Errorf("sortAsc(%+v) returned %+v, want %+v", testCase.data, res, testCase.answer)
		}
	}
	tests = []TestCases{
		{data: [8]int{1, 2, 3, 4, 5, 8, 7, 6}, answer: [8]int{8, 7, 6, 5, 4, 3, 2, 1}},
		{data: [8]int{1, 2, 3, 4, 5, 6, 7, 8}, answer: [8]int{8, 7, 6, 5, 4, 3, 2, 1}},
	}
	for _, testCase := range tests {
		res := sortDesc(testCase.data)
		if res != testCase.answer {
			t.Errorf("sortDesc(%+v) returned %+v, want %+v", testCase.data, res, testCase.answer)
		}
	}
}
